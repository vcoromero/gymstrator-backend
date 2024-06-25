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
