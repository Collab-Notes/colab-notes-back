package controllers

import (
	"net/http"
	"strconv"

	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/dto"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/Collab-Notes/colab-notes-back/repository"
	"github.com/gin-gonic/gin"
)

// UpdateVaultRoleHandler actualiza el rol de un usuario dentro de un vault.
//
// @Summary Actualizar el rol de un usuario en un vault
// @Description Permite cambiar el rol de un colaborador existente en un vault (por ejemplo, de viewer a collaborator).
// @Tags Vaults
// @Accept json
// @Produce json
// @Param id path int true "ID del Vault"
// @Param request body dto.UpdateVaultRoleRequest true "Datos para actualizar el rol"
// @Success 200 {object} dto.MessageResponse
// @Failure 400,404,403,500 {object} dto.ErrorResponse
// @Router /vaults/{id}/role [patch]
func UpdateVaultRoleHandler(c *gin.Context) {
	vaultIDParam := c.Param("id")
	vaultID, err := strconv.Atoi(vaultIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de vault inválido"})
		return
	}

	var requestData dto.UpdateVaultRoleRequest
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
