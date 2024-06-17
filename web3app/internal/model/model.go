package model

type Response struct {
	Success bool        `json:"succes"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
