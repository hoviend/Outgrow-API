package dto

type Err struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
