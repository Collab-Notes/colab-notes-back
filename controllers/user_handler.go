package controllers

import (
	"net/http"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/gin-gonic/gin"
)

// SearchUsersHandler maneja la búsqueda de usuarios por nombre parcial.
func SearchUsersHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'q' es requerido"})
		return
	}

	var users []models.User
	if err := common.DB.Where("name ILIKE ?", "%"+query+"%").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuarios"})
		return
	}

	c.JSON(http.StatusOK, users)
}
