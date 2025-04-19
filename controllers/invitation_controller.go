package controllers

import (
	"net/http"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/Collab-Notes/colab-notes-back/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// InviteUserToVaultHandler maneja la invitación de un usuario a un vault.
// URL: POST /vaults/:id/invite
func InviteUserToVaultHandler(c *gin.Context) {
	vaultIDStr := c.Param("id")
	vaultID, err := uuid.Parse(vaultIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de vault inválido"})
		return
	}

	var requestData struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil || requestData.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El campo username es requerido"})
		return
	}

	var user models.User
	if err := common.DB.Where("name = ?", requestData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Verificar si ya tiene un permiso en el vault
	if _, err := repository.GetVaultPermission(vaultID, user.ID); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El usuario ya pertenece al vault"})
		return
	}

	// Crear el permiso directamente como "collaborator"
	newPermission := models.VaultPermission{
		VaultID:     vaultID,
		UserID:      user.ID,
		AccessLevel: "collaborator",
	}
	if err := repository.CreateVaultPermission(&newPermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar colaborador"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Colaborador agregado exitosamente"})
}
