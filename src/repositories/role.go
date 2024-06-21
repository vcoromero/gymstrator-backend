package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/vcoromero/gymstration-backend/database"
	"github.com/vcoromero/gymstration-backend/src/models"
)

func FetchRoles(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role
	query := "SELECT id, name, created_at, updated_at, deleted_at FROM roles ORDER BY created_at ASC"
	rows, err := database.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var role models.Role
		if err := rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return roles, nil
}

func GetRoleById(ctx context.Context, id string) (models.Role, error) {
	var role models.Role

	query := "SELECT id, name, created_at, updated_at, deleted_at FROM roles WHERE id = ?"
	row := database.DB.QueryRowContext(ctx, query, id)

	if err := row.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt, &role.DeletedAt); err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func CreateRole(ctx context.Context, name string) (models.Role, error) {
	var role models.Role

	id := uuid.New().String()

	query := "INSERT INTO roles (id, name) VALUES (?, ?)"
	_, err := database.DB.ExecContext(ctx, query, id, name)
	if err != nil {
		return models.Role{}, err
	}

	role, err = GetRoleById(ctx, id)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func UpdateRole(ctx context.Context, id string, name string) (models.Role, error) {
	var role models.Role

	query := "UPDATE roles SET name = ? WHERE id = ?"
	_, err := database.DB.ExecContext(ctx, query, name, id)
	if err != nil {
		return models.Role{}, err
	}

	role, err = GetRoleById(ctx, id)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}
