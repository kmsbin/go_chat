package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	migration "go_chat/migrations"
	_ "go_chat/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	uuid "github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
var hub Hub = newHub()

func init() {
	_ = godotenv.Load(".env")
	db, err := gorm.Open(postgres.Open(os.Getenv("DBURI")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	mg := migration.Migration{GormConection: db}

	mg.Up()
	log.Println(db)
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	go hub.Run()

	r.Get("/enter-room", func(w http.ResponseWriter, r *http.Request) {
		uuidToken := uuid.NewV4()
		w.Write([]byte(uuidToken.String()))

	})
	r.Get("/getAllUsers", getAllUsers)
	r.Get("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		serveWs(&hub, w, r)
	})

	http.ListenAndServe(string(":"+os.Getenv("PORT")), r)
}

type ActiveUsers struct {
	Clients []Client `json:"clients"`
}

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users := ActiveUsers{Clients: make([]Client, 0)}

	for _, client := range hub.Registered {
		log.Println(client)
		users.Clients = append(users.Clients, client)
	}
	activeUsersJson, _ := json.Marshal(users)
	log.Println(activeUsersJson)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Efsdfhsdfh", "application/json")
	w.Write(activeUsersJson)
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	socket, _ := upgrader.Upgrade(w, r, nil)
	client := newClient(socket, id)
	go client.ReadMessagePool()
	go client.WriteMessagePool()
	log.Println(id)

	hub.Register(client)
}
