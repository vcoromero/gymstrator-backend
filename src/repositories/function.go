package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
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

func GetFunction(ctx context.Context, id string) (models.Function, error) {
	var function models.Function

	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM functions WHERE id = ?"
	row := database.DB.QueryRowContext(ctx, query, id)

	if err := row.Scan(&function.ID, &function.Name, &function.Description, &function.CreatedAt, &function.UpdatedAt, &function.DeletedAt); err != nil {
		return models.Function{}, nil
	}

	return function, nil
}

func CreateFunction(ctx context.Context, name string, description string) (models.Function, error) {
	id := uuid.New().String()

	query := "INSERT INTO functions (id, name, description) VALUES (?, ?,?)"
	_, err := database.DB.ExecContext(ctx, query, id, name, description)
	if err != nil {
		return models.Function{}, err
	}

	function, err := GetFunction(ctx, id)
	if err != nil {
		return models.Function{}, err
	}

	return function, nil
}

func UpdateFunction(ctx context.Context, id string, name string, description string) (models.Function, error) {
	query := "UPDATE functions SET name = ?, description = ? WHERE id = ?"
	_, err := database.DB.ExecContext(ctx, query, name, description, id)

	if err != nil {
		return models.Function{}, err
	}

	function, err := GetFunction(ctx, id)
	if err != nil {
		return models.Function{}, err
	}

	return function, nil
}

func DeleteFunction(ctx context.Context, id string) error {
	deleteAt := time.Now()
	query := "UPDATE functions SET deleted_at = ? WHERE id = ?"

	_, err := database.DB.ExecContext(ctx, query, deleteAt, id)
	if err != nil {
		return err
	}

	return nil
}
