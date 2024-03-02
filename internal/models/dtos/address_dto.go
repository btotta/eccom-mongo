package dtos

import "eccom-mongo/internal/models"

type AddressRegisterRequest struct {
	Street       string `json:"street" binding:"required"`
	Number       string `json:"number" binding:"required"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	ZipCode      string `json:"zip_code" binding:"required"`
}

func (a *AddressRegisterRequest) ToAddress() *models.Address {
	return &models.Address{
		Street:       a.Street,
		Number:       a.Number,
		Complement:   a.Complement,
		Neighborhood: a.Neighborhood,
		City:         a.City,
		State:        a.State,
		Country:      a.Country,
		ZipCode:      a.ZipCode,
	}
}
