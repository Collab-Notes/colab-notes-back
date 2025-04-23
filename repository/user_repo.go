package repository

import (
	"github.com/Collab-Notes/colab-notes-back/common"
	"github.com/Collab-Notes/colab-notes-back/models"
)

func SearchUsers(query string, limit int) ([]models.User, error) {
	var users []models.User

	if err := common.DB.Where("Name ILIKE ?", "%"+query+"%").Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
