package controller

import (
	"eccom-mongo/internal/database"
	"eccom-mongo/internal/models/dtos"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	CreateProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	SearchProduct(c *gin.Context)
}

type productHandler struct {
	productDAO database.ProductDAOInterface
}

func NewProductController(productDAO database.ProductDAOInterface) ProductHandler {
	return &productHandler{
		productDAO: productDAO,
	}
}

// @Summary Create a product
// @Description Create a product
// @Tags product
// @Accept json
// @Produce json
// @Param product body dtos.ProductDTO true "Product"
// @Success 201 {object} dtos.ProductDTO
// @Router /product [post]
func (ph *productHandler) CreateProduct(c *gin.Context) {

	//TODO: Implement User roles, so only admins can create a product

	var productDTO dtos.ProductDTO
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := ph.productDAO.CreateProduct(c, productDTO.ToModel())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, dtos.NewProductDTO(createdProduct))

}

// @Summary Get a product
// @Description Get a product
// @Tags product
// @Accept json
// @Produce json
// @Param sku path string true "Product SKU"
// @Success 200 {object} dtos.ProductDTO
// @Router /product/{sku} [get]
func (ph *productHandler) GetProduct(c *gin.Context) {
	productID := c.Param("sku")

	product, err := ph.productDAO.GetProduct(c, productID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dtos.NewProductDTO(product))
}

// @Summary Search a product
// @Description Search a product
// @Tags product
// @Accept json
// @Produce json
// @Param terms query string true "Search terms"
// @Success 200 {object} []dtos.ProductDTO
// @Router /product/{terms} [get]
func (ph *productHandler) SearchProduct(c *gin.Context) {

	terms := c.Query("terms")

	if terms == "" {
		c.JSON(400, gin.H{"error": "search terms are required"})
		return
	}

	keys := strings.Split(strings.ToLower(terms), ",")

	products, err := ph.productDAO.SearchProduct(c, keys)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)

}
