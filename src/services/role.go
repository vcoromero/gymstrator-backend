package services

import (
	"github.com/vcoromero/gymstration-backend/src/models"
	"github.com/vcoromero/gymstration-backend/src/repositories"
)

func GetRoles() ([]models.Role, error) {
	return repositories.FetchRoles()
}
