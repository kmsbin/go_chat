package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	uuid "github.com/satori/go.uuid"
)

func router() *chi.Mux {
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
	return r
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
	w.Write(activeUsersJson)
}