package controllers

import (
	"net/http"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/gin-gonic/gin"
)

// SearchUsersHandler maneja la búsqueda de usuarios por nombre parcial.
// @Summary Buscar usuarios
// @Description Busca usuarios cuyo nombre coincida parcialmente con el query enviado.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param q query string true "Texto de búsqueda parcial"
// @Success 200 {array} models.User
// @Failure 400 {object} dto.ErrorResponse "Parámetro 'q' faltante"
// @Failure 500 {object} dto.ErrorResponse "Error al buscar en la base de datos"
// @Router /users/search [get]
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
