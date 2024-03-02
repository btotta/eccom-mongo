package dtos

import "eccom-mongo/internal/models"

type AddressDTO struct {
	Id           string `json:"id"`
	Street       string `json:"street" binding:"required"`
	Number       string `json:"number" binding:"required"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
	ZipCode      string `json:"zip_code" binding:"required"`
	MainAddress  bool   `json:"main_address"`
}

func (a *AddressDTO) ToAddress() *models.Address {
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

func NewAddressDTO(address *models.Address) *AddressDTO {
	return &AddressDTO{
		Id:           address.ID.Hex(),
		Street:       address.Street,
		Number:       address.Number,
		Complement:   address.Complement,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		ZipCode:      address.ZipCode,
		MainAddress:  address.MainAddress,
	}
}
