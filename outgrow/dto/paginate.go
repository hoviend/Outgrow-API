package dto

type PaginateParam struct {
	Page    int `schema:"page" json:"page,omitempty" validate:"omitempty"`
	PerPage int `schema:"per_page" json:"per_page,omitempty" validate:"omitempty"`
}

type PaginateResponse struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

func (p PaginateParam) ToResponse(total int) PaginateResponse {
	return PaginateResponse{
		Page:    p.Page,
		PerPage: p.PerPage,
		Total:   total,
	}
}

func (p PaginateParam) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func (p PaginateParam) MaxPageByTotal(total int) int {
	return (total + p.PerPage - 1) / p.PerPage
}

func (p *PaginateParam) AdjustByTotal(total int) {
	if p.PerPage < 1 {
		p.PerPage = 10
	} else if p.PerPage > 100 {
		p.PerPage = 100
	}

	maxPage := p.MaxPageByTotal(total)
	if p.Page > maxPage {
		p.Page = maxPage
	}

	if p.Page < 1 {
		p.Page = 1
	}
}
