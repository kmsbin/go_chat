package main

import (
	"log"
	"net/http"
	"os"

	_ "go_chat/entities"

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
	// mg := migration.Migration{GormConection: db}
	// mg.Up()
	// mg.Seed()
	log.Println(db)
}

func main() {

	http.ListenAndServe(string(":"+os.Getenv("PORT")), router())
}
