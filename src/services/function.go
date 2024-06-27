package services

import (
	"context"

	"github.com/vcoromero/gymstration-backend/src/models"
	"github.com/vcoromero/gymstration-backend/src/repositories"
)

func GetFunctions(ctx context.Context) ([]models.Function, error) {
	functions, err := repositories.GetFunctions(ctx)
	if err != nil {
		return nil, err
	}
	return functions, nil
}

func GetFunction(ctx context.Context, id string) (models.Function, error) {
	function, err := repositories.GetFunction(ctx, id)
	if err != nil {
		return models.Function{}, nil
	}

	return function, nil
}

func CreateFunction(ctx context.Context, name string, description string) (models.Function, error) {
	function, err := repositories.CreateFunction(ctx, name, description)
	if err != nil {
		return models.Function{}, err
	}

	return function, nil
}

func UpdateFunction(ctx context.Context, id string, name string, description string) (models.Function, error) {
	function, err := repositories.UpdateFunction(ctx, id, name, description)
	if err != nil {
		return models.Function{}, err
	}

	return function, nil
}

func DeleteFunction(ctx context.Context, id string) error {
	err := repositories.DeleteFunction(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
