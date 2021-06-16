package main

import (
	"log"
	"net/http"
	"os"

	migration "go_chat/migrations"
	_ "go_chat/model"

	_ "github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	_ = godotenv.Load(".env")
	db, err := gorm.Open(postgres.Open(os.Getenv("DBURI")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	mg := migration.Migration{GormConection: db}

	http.ListenAndServe(string(":"+os.Getenv("PORT")), router())
	mg.Up()
	log.Println(db)
}

func main() {

	http.ListenAndServe(string(":"+os.Getenv("PORT")), router())
}
