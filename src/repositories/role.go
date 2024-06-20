package repositories

import (
	"github.com/vcoromero/gymstration-backend/database"
	"github.com/vcoromero/gymstration-backend/src/models"
)

func FetchRoles() ([]models.Role, error) {
	var roles []models.Role
	rows, err := database.DB.Query("SELECT id, name, created_at, updated_at, deleted_at FROM roles")
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
