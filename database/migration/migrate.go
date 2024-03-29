package migration

import (
	"fmt"
	"go-enamad-test/database"
	"go-enamad-test/models"
	"log"
)

func Migrate() {

	db := database.Connection()

	err := db.AutoMigrate(&models.Company{})
	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration Done ..")
}
