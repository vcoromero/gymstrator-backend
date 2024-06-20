package models

import (
	"database/sql"
	"encoding/json"
)

type Role struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

func (r Role) MarshalJSON() ([]byte, error) {
	type Alias Role
	return json.Marshal(&struct {
		DeletedAt interface{} `json:"deleted_at"`
		*Alias
	}{
		Alias:     (*Alias)(&r),
		DeletedAt: ifNullString(r.DeletedAt),
	})
}

func ifNullString(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}
