package main

import (
	"fmt"
	"log"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
)

func main() {
	config := common.LoadConfig()
	common.ConnectDB(config)

	err := common.DB.AutoMigrate(
		&models.User{},
		&models.Vault{},
		&models.Note{},
		&models.VaultPermission{},
		&models.NotePermission{},
		&models.Tag{},
		&models.NoteTag{},
		&models.NoteAttachment{},
	)
	if err != nil {
		log.Fatalf("Error en la migración de la base de datos: %v", err)
	}

	fmt.Println("Migración completada exitosamente.")
}
