package controller

import (
	"eccom-mongo/internal/models/dtos"
	"eccom-mongo/internal/database"
	"eccom-mongo/internal/utils"

	"github.com/gin-gonic/gin"
)

type AddressHandler interface {
	CreateAddress(c *gin.Context)
	// DeleteAddress(c *gin.Context)
	// GetAddress(c *gin.Context)
}

type addressHandler struct {
	addressDAO database.AddressDAOInterface
	userDAO    database.UserDAOInterface
}

func NewAddressController(addressDAO database.AddressDAOInterface, userDAO database.UserDAOInterface) AddressHandler {
	return &addressHandler{
		addressDAO: addressDAO,
		userDAO:    userDAO,
	}
}

// @Summary Create an address
// @Description Create an address
// @Tags address
// @Accept json
// @Produce json
// @Param address body dtos.AddressRegisterRequest true "Address to be created"
// @Success 200 {object} models.Address
// @Router /address [post]
func (a *addressHandler) CreateAddress(c *gin.Context) {

	user := utils.ExtractUserFromRequest(c)

	var addressRequest dtos.AddressRegisterRequest

	if err := c.ShouldBindJSON(&addressRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	address := addressRequest.ToAddress()
	err := a.addressDAO.CreateAddress(c, address)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = a.userDAO.AddAddress(c, user.Email, address.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, address)
}
