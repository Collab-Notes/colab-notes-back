package controllers

import (
	"net/http"

	"time"

	"github.com/Collab-Notes/colab-notes-back/models"

	"github.com/gin-gonic/gin"
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateNotesRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

// POST (TO CREATE A NEW VAULT)
func CreateNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateNotesRequest

		// Validar el JSON de entrada
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generar un OwnerID único
		VaultID := c.Param("id")

		// Verificar si el vault existe
		var vault models.Vault
		if err := db.First(&vault, "id = ?", VaultID).Error; err != nil {

			//devolver error si no encuentra el vault
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Vault no encontrado"})
			return

		}

		//verificación de permisos
		if vault.OwnerID.String() != req.UserID {

			var permission models.VaultPermission
			if err := db.First(&permission, "vault_id = ? AND user_id = ?", vault.ID, req.UserID).Error; err != nil {
				// Si no existe, devolver error
				c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para acceder a este vault"})
				return
			}

		}

		//para la creación de las notas
		note := models.Note{
			ID:        vault.ID,
			VaultID:   vault.ID,
			OwnerID:   vault.OwnerID,
			Title:     req.Title,
			CreatedAt: time.Now(),
		}

		// Guardar en la base de datos
		if err := db.Create(&note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la nota"})
			return
		}

		// Responder con el Vault creado
		c.JSON(http.StatusCreated, gin.H{
			"id":      note.ID,
			"message": "note created successfully",
		})
	}
}

// GET (TO GET ALL VAULTS)
func GetNote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var Note []models.Vault

		// Obtener todos los Vaults de la base de datos
		if err := db.Find(&Note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las notas"})
			return
		}

		// Responder con la lista de Vaults
		c.JSON(http.StatusOK, Note)
	}
}
