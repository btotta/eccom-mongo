package controller

import (
	"eccom-mongo/internal/database"
	"eccom-mongo/internal/models/dtos"
	"eccom-mongo/internal/utils"

	"github.com/gin-gonic/gin"
)

type AddressHandler interface {
	CreateAddress(c *gin.Context)
	DeleteAddress(c *gin.Context)
	GetAddress(c *gin.Context)
	GetAllAddress(c *gin.Context)
	MarkAddressAsMain(c *gin.Context)
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
// @Param address body dtos.AddressDTO true "Address to be created"
// @Success 200 {object} models.Address
// @Router /address [post]
func (a *addressHandler) CreateAddress(c *gin.Context) {

	user := utils.ExtractUserFromRequest(c)

	var addressRequest dtos.AddressDTO

	if err := c.ShouldBindJSON(&addressRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	address := addressRequest.ToAddress()
	address.UserID = user.ID

	createdAddress, err := a.addressDAO.CreateAddress(c, address)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dtos.NewAddressDTO(createdAddress))
}

// @Summary Delete an address
// @Description Delete an address
// @Tags address
// @Accept json
// @Produce json
// @Param id path string true "Address ID"
// @Success 200 {string} string "Address deleted"
// @Router /address/{id} [delete]
// @Security ApiKeyAuth
func (a *addressHandler) DeleteAddress(c *gin.Context) {

	addressID := c.Param("id")

	err := a.addressDAO.DeleteAddress(c, addressID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Address deleted"})
}

// @Summary Get an address
// @Description Get an address
// @Tags address
// @Accept json
// @Produce json
// @Param id path string true "Address ID"
// @Success 200 {object} dtos.AddressDTO
// @Router /address/{id} [get]
// @Security ApiKeyAuth
func (a *addressHandler) GetAddress(c *gin.Context) {
	addressID := c.Param("id")

	address, err := a.addressDAO.GetAddress(c, addressID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dtos.NewAddressDTO(address))
}

// @Summary Get all addresses
// @Description Get all addresses
// @Tags address
// @Accept json
// @Produce json
// @Success 200 {array} dtos.AddressDTO
// @Router /address [get]
// @Security ApiKeyAuth
func (a *addressHandler) GetAllAddress(c *gin.Context) {

	user := utils.ExtractUserFromRequest(c)

	addresses, err := a.addressDAO.GetAllAddressByUserID(c, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var resp []dtos.AddressDTO

	for _, address := range addresses {
		resp = append(resp, *dtos.NewAddressDTO(&address))
	}

	c.JSON(200, resp)

}

// @Summary Mark an address as default
// @Description Mark an address as default
// @Tags address
// @Accept json
// @Produce json
// @Param id path string true "Address ID"
// @Success 200 {string} dtos.AddressDTO
// @Router /address/{id}/main [put]
// @Security ApiKeyAuth
func (a *addressHandler) MarkAddressAsMain(c *gin.Context) {
	addressID := c.Param("id")

	user := utils.ExtractUserFromRequest(c)

	address, err := a.addressDAO.MarkAddressAsMain(c, addressID, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dtos.NewAddressDTO(address))

}
