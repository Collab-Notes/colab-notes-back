package main

import (
	"fmt"
	"log"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/controllers"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración y conectar a la base de datos
	config := common.LoadConfig()
	common.ConnectDB(config)

	// Migrar modelos
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

	//enrutador mediante gin
	router := gin.Default()

	// Endpoint 1: Buscar usuarios
	router.GET("/users/search", controllers.SearchUsersHandler)

	// Endpoint 3: Invitar a un usuario a un vault
	router.POST("/vaults/:id/invite", controllers.InviteUserToVaultHandler)

	// Endpoint 5:
	router.PATCH("/vaults/:id/role", controllers.UpdateVaultRoleHandler)

	// Rutas
	router.POST("/vaults", controllers.CreateVault())
	router.GET("/vaults/:id", controllers.GetVault())
	router.POST("/vaults/:id/notes", controllers.CreateNote())

	// Iniciar servidor
	router.Run(":8080")

}
