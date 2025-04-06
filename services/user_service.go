package services

import (
	"github.com/Collab-Notes/colab-notes-back/repository"

	"github.com/Collab-Notes/colab-notes-back/models"
)

// SearchUsersService busca usuarios con base en un query parcial.
func SearchUsersService(query string) ([]models.User, error) {
	const limit = 10
	return repository.SearchUsers(query, limit)
}
