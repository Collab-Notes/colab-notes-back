package repository

import (
	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/google/uuid"
)

// SearchUsers busca usuarios cuyo username coincida parcialmente con el query.
func SearchUsers(query string, limit int) ([]models.User, error) {
	var users []models.User

	// Realiza la consulta con LIKE para buscar coincidencias parciales
	if err := common.DB.Where("Name ILIKE ?", "%"+query+"%").Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// AcceptVaultInvitation actualiza el permiso de "invited" a "collaborator" para indicar que la invitación fue aceptada.
func AcceptVaultInvitation(vaultID uint, userID uuid.UUID) error {
	var vp models.VaultPermission
	err := common.DB.Where("vault_id = ? AND user_id = ? AND access_level = ?", vaultID, userID, "invited").First(&vp).Error
	if err != nil {
		return err
	}
	vp.AccessLevel = "collaborator"
	return common.DB.Save(&vp).Error
}
