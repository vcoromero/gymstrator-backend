package models

type ResponseAPI struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
