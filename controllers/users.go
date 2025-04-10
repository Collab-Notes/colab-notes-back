package controllers

import (
	"net/http"
	"time"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/gin-gonic/gin"
)

// GET (TO GET ALL VAULTS)
func GetUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")

		// Obtener información del usuario
		var user models.User
		if err := common.DB.First(&user, "id = ?", userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Obtener los Vaults donde el usuario es propietario o tiene permisos
		var vaults []models.Vault
		if err := common.DB.Preload("Notes").Preload("Permissions").Where("owner_id = ?", userID).Or("id IN (?)",
			common.DB.Table("vault_permissions").Select("vault_id").Where("user_id = ?", userID),
		).Find(&vaults).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los Vaults"})
			return
		}

		// Construir la respuesta
		response := gin.H{
			"user": gin.H{
				"id":        user.ID,
				"name":      user.Name,
				"email":     user.Email,
				"createdAt": user.CreatedAt.Format(time.RFC3339),
			},
			"vaults": vaults,
		}

		// Responder con los datos del usuario y sus Vaults
		c.JSON(http.StatusOK, response)
	}
}
