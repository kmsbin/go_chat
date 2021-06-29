package main

import (
	"encoding/json"
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

	r.Get("/login", GenerateToken)
	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/enter-room", func(w http.ResponseWriter, r *http.Request) {
			uuidToken := uuid.NewV4()
			response, _ := json.Marshal(struct {
				Key string `json:"key"`
			}{uuidToken.String()})

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))

		})
		r.Get("/getAllUsers", getAllUsers)
		r.Get("/getAllUsersById/{id}", getAllUsersById)
		r.Get("/ws/{username}/{id}", func(w http.ResponseWriter, r *http.Request) {
			serveWs(&hub, w, r)
		})
	})
	return r
}

type ActiveUsers struct {
	Clients []Client `json:"clients"`
}

func getAllUsers(w http.ResponseWriter, _ *http.Request) {
	users := ActiveUsers{Clients: make([]Client, 0)}

	for _, client := range hub.Registered {
		users.Clients = append(users.Clients, client)
	}
	activeUsersJson, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(activeUsersJson)
}
func getAllUsersById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	users := ActiveUsers{Clients: make([]Client, 0)}

	for _, client := range hub.Registered {
		if client.Id != id {
			users.Clients = append(users.Clients, client)
		}
	}
	activeUsersJson, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(activeUsersJson)
}
