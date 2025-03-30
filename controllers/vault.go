package controllers

import (
	"net/http"

	"time"

	"github.com/Collab-Notes/colab-notes-back/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateVaultRequest struct {
	Name          string             `json:"name" binding:"required"`
	Description   string             `json:"description" binding:"required"`
	Collaborators []CollaboratorRole `json:"collaborators,omitempty"`
}

type CollaboratorRole struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
	Role   string    `json:"role" binding:"required"`
}

// POST (TO CREATE A NEW VAULT)
func CreateVault(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateVaultRequest

		// Validar el JSON de entrada
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generar un OwnerID único
		ownerID := uuid.New()

		// Verificar si el OwnerID existe en la tabla "users"
		var owner models.User
		if err := db.First(&owner, "id = ?", ownerID).Error; err != nil {
			// Si no existe, crear un usuario temporal
			newUser := models.User{
				ID:        ownerID,
				Name:      "kkk",
				Email:     ownerID.String() + "@example.com",
				CreatedAt: time.Now(),
			}
			if err := db.Create(&newUser).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
				return
			}
		}

		// Crear el nuevo Vault
		vault := models.Vault{
			OwnerID:     ownerID,
			Name:        req.Name,
			Description: req.Description,
			CreatedAt:   time.Now(),
		}

		// Guardar en la base de datos
		if err := db.Create(&vault).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el Vault"})
			return
		}

		//Admin
		adminPermission := models.VaultPermission{
			UserID:      ownerID,
			VaultID:     vault.ID,
			AccessLevel: "admin",
		}

		if err := db.Create(&adminPermission).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al asignar el rol de admin al creador"})
			return
		}

		// Agregar colaboradores
		for _, collaborator := range req.Collaborators {
			// Verificar si el colaborador existe en la tabla "users"
			var user models.User
			if err := db.First(&user, "id = ?", collaborator.UserID).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "El colaborador con ID " + collaborator.UserID.String() + " no existe"})
				return
			}

			// Crear el permiso para el colaborador
			permission := models.VaultPermission{
				UserID:      collaborator.UserID,
				VaultID:     vault.ID,
				AccessLevel: collaborator.Role,
			}
			if err := db.Create(&permission).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar el colaborador con ID " + collaborator.UserID.String()})
				return
			}
		}

		// Responder con el Vault creado
		c.JSON(http.StatusCreated, gin.H{
			"id":      vault.ID,
			"message": "Vault created successfully",
		})
	}
}

// GET (TO GET ALL VAULTS)
func GetVault(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var vaults []models.Vault

		// Obtener todos los Vaults de la base de datos
		if err := db.Find(&vaults).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los Vaults"})
			return
		}

		// Responder con la lista de Vaults
		c.JSON(http.StatusOK, vaults)
	}
}
