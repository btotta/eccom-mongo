package dtos

import "eccom-mongo/internal/models"

type ProductDTO struct {
	ID             string   `json:"id"`
	Name           string   `json:"name" binding:"required"`
	Description    string   `json:"description" binding:"required"`
	Sku            string   `json:"sku" binding:"required"`
	Price          float64  `json:"price" `
	Stock          int      `json:"quantity" binding:"required"`
	Active         bool     `json:"active"`
	NameUrl        string   `json:"name_url" binding:"required"`
	Categorys      []string `json:"categorys"`
	Brands         []string `json:"brands" `
	Images         []string `json:"images" `
	PrincipalImage string   `json:"principal_image" binding:"required"`
}

func (dto *ProductDTO) ToModel() *models.Product {
	return &models.Product{
		ID:             dto.ID,
		Name:           dto.Name,
		Description:    dto.Description,
		Price:          dto.Price,
		Sku:            dto.Sku,
		Stock:          dto.Stock,
		Active:         dto.Active,
		NameUrl:        dto.NameUrl,
		Categorys:      dto.Categorys,
		Brands:         dto.Brands,
		Images:         dto.Images,
		PrincipalImage: dto.PrincipalImage,
	}
}

func NewProductDTO(product *models.Product) *ProductDTO {
	return &ProductDTO{
		ID:             product.ID,
		Name:           product.Name,
		Description:    product.Description,
		Price:          product.Price,
		Sku:            product.Sku,
		Stock:          product.Stock,
		Active:         product.Active,
		NameUrl:        product.NameUrl,
		Categorys:      product.Categorys,
		Brands:         product.Brands,
		Images:         product.Images,
		PrincipalImage: product.PrincipalImage,
	}
}
