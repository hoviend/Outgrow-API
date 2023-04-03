package dto

type SuccessResponse struct {
	Message    string            `json:"message,omitempty"`
	Data       interface{}       `json:"data"`
	Pagination *PaginateResponse `json:"pagination,omitempty"`
}
