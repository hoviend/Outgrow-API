package dto

import "github.com/google/uuid"

type SuccessResponse struct {
	Message    string            `json:"message,omitempty"`
	Data       interface{}       `json:"data"`
	Pagination *PaginateResponse `json:"pagination,omitempty"`
}

type DefaultCreateIDResponse struct {
	ID int `json:"id"`
}

type DefaultCreateUUIDResponse struct {
	ID uuid.UUID `json:"id"`
}
