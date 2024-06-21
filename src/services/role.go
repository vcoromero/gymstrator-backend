package services

import (
	"context"

	"github.com/vcoromero/gymstration-backend/src/models"
	"github.com/vcoromero/gymstration-backend/src/repositories"
)

func GetRoles(ctx context.Context) ([]models.Role, error) {
	roles, err := repositories.FetchRoles(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRole(ctx context.Context, id string) (models.Role, error) {
	role, err := repositories.GetRole(ctx, id)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func CreateRole(ctx context.Context, name string) (models.Role, error) {
	role, err := repositories.CreateRole(ctx, name)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func UpdateRole(ctx context.Context, id string, name string) (models.Role, error) {
	role, err := repositories.UpdateRole(ctx, id, name)
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func DeleteRole(ctx context.Context, id string) error {
	err := repositories.DeleteRole(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
