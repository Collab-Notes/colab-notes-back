package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/Collab-Notes/colab-notes-back/repository"
)

func UpdateVaultRoleHandler(c *gin.Context) {
	vaultIDParam := c.Param("id")
	vaultID, err := strconv.Atoi(vaultIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de vault inválido"})
		return
	}

	var requestData struct {
		Username string `json:"username"`
		NewRole  string `json:"new_role"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar el usuario
	var user models.User
	if err := common.DB.Where("name = ?", requestData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Actualizar el rol
	if err := repository.UpdateVaultPermission(uint(vaultID), user.ID, requestData.NewRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el rol"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Rol actualizado exitosamente"})
}
