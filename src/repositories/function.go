package repositories

import (
	"context"

	"github.com/vcoromero/gymstration-backend/database"
	"github.com/vcoromero/gymstration-backend/src/models"
)

func GetFunctions(ctx context.Context) ([]models.Function, error) {
	var functions []models.Function

	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM functions ORDER BY created_at ASC"
	rows, err := database.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var function models.Function
		if err := rows.Scan(&function.ID, &function.Name, &function.Description, &function.CreatedAt, &function.UpdatedAt, &function.DeletedAt); err != nil {
			return nil, err
		}
		functions = append(functions, function)
	}

	return functions, nil
}
