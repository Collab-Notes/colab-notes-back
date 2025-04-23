package controllers

import (
	"net/http"
	"time"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateNotesRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

// CreateNote (endpoint)
// @Summary Crear una nueva nota
// @Description Crea una nueva nota dentro de un vault específico
// @Tags Notes
// @Accept json
// @Produce json
// @Param id path string true "ID del Vault" example("123e4567-e89b-12d3-a456-426614174000")
// @Param request body CreateNotesRequest true "Datos de la nota"
// @Success 201 {object} map[string]interface{} "Nota creada"
// @Failure 400 {object} map[string]string "Error en la solicitud"
// @Failure 403 {object} map[string]string "Permisos insuficientes"
// @Failure 404 {object} map[string]string "Vault no encontrado"
// @Failure 500 {object} map[string]string "Error interno"
// @Router /vaults/{id}/notes [post]
func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateNotesRequest

		// Validar el JSON de entrada
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Obtener el ID del Vault desde los parámetros de la URL
		VaultID := c.Param("id")

		// Verificar si el Vault existe
		var vault models.Vault
		if err := common.DB.First(&vault, "id = ?", VaultID).Error; err != nil {
			// Devolver error si no encuentra el Vault
			c.JSON(http.StatusNotFound, gin.H{"error": "Vault no encontrado"})
			return
		}

		// Verificación de permisos
		if vault.OwnerID.String() != req.UserID {
			var permission models.VaultPermission
			if err := common.DB.First(&permission, "vault_id = ? AND user_id = ?", vault.ID, req.UserID).Error; err != nil {
				// Si no existe, devolver error
				c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para acceder a este Vault"})
				return
			}
		}

		// Crear la nueva nota
		note := models.Note{
			ID:        vault.ID,
			VaultID:   vault.ID,
			OwnerID:   uuid.MustParse(req.UserID),
			Title:     req.Title,
			CreatedAt: time.Now(),
		}

		// Guardar en la base de datos
		if err := common.DB.Create(&note).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la nota"})
			return
		}

		// Responder con la nota creada
		c.JSON(http.StatusCreated, gin.H{
			"id":      note.ID,
			"message": "Note created successfully",
		})
	}
}

// GetNote (endpoint)
// @Summary Obtener todas las notas
// @Description Devuelve una lista de todas las notas existentes
// @Tags Notes
// @Produce json
// @Success 200 {array} models.Note "Lista de notas"
// @Failure 500 {object} map[string]string "Error interno"
// @Router /notes [get]
func GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var notes []models.Note

		// Obtener todas las notas de la base de datos
		if err := common.DB.Find(&notes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las notas"})
			return
		}

		// Responder con la lista de notas
		c.JSON(http.StatusOK, notes)
	}
}
