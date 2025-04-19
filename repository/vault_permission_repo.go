package repository

import (
	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
	"github.com/google/uuid"
)

// GetVaultPermission busca si ya existe un permiso para el usuario en el vault.
func GetVaultPermission(vaultID uuid.UUID, userID uuid.UUID) (*models.VaultPermission, error) {
	var vp models.VaultPermission
	err := common.DB.Where("vault_id = ? AND user_id = ?", vaultID, userID).First(&vp).Error
	return &vp, err
}

// CreateVaultPermission guarda un nuevo permiso en la base de datos.
func CreateVaultPermission(vp *models.VaultPermission) error {
	return common.DB.Create(vp).Error
}

// UpdateVaultPermission actualiza el rol (AccessLevel) de un usuario en un vault.
func UpdateVaultPermission(vaultID uint, userID uuid.UUID, newRole string) error {
	var vp models.VaultPermission
	if err := common.DB.Where("vault_id = ? AND user_id = ?", vaultID, userID).First(&vp).Error; err != nil {
		return err
	}

	vp.AccessLevel = newRole
	return common.DB.Save(&vp).Error
}
