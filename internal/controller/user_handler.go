package controller

import (
	"eccom-mongo/internal/database"
	"eccom-mongo/internal/middleware"
	"eccom-mongo/internal/models"
	"eccom-mongo/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	// EditUser(c *gin.Context)
	// ChangePasswordUser(c *gin.Context)
	// DeleteUser(c *gin.Context)
	// GetUser(c *gin.Context)
}

type userHandler struct {
	userDAO database.UserDAOInterface
}

func NewUserController(userDAO database.UserDAOInterface) UserHandler {
	return &userHandler{
		userDAO: userDAO,
	}
}

type userRegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Document        string `json:"document" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type userLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type userLoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// @Summary Create user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userRegisterRequest true "User data"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /user [post]
func (uh *userHandler) CreateUser(c *gin.Context) {
	var userCreate userRegisterRequest
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userCreate.Password != userCreate.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this password cannot be used"})
		return
	}

	if !utils.ValidatePassword(userCreate.Password, userCreate.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this password cannot be used"})
		return
	}

	if !utils.ValidateEmail(userCreate.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this email cannot be used"})
		return
	}

	passwordHash, err := utils.HashBcrypt(userCreate.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this password cannot be used"})
		return
	}

	user, err := uh.userDAO.FindByEmail(c, userCreate.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this email cannot be used"})
		return
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this email cannot be used"})
		return
	}

	user, err = uh.userDAO.FindByCPF(c, userCreate.Document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this document cannot be used"})
		return
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this document cannot be used"})
		return
	}

	newUser := models.User{
		Name:         userCreate.Name,
		LastName:     userCreate.LastName,
		Document:     userCreate.Document,
		Email:        userCreate.Email,
		PasswordHash: passwordHash,
	}

	err = uh.userDAO.CreateUser(c, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

// @Summary Login user
// @Description Login user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userLoginRequest true "User data"
// @Success 200 {object} userLoginResponse
// @Failure 400 {object} string
// @Router /user/login [post]
func (uh *userHandler) LoginUser(c *gin.Context) {
	var userLogin userLoginRequest
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userDAO.FindByEmail(c, userLogin.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userLogin.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login credentials"})
		return
	}

	token, err := middleware.NewAuthenticationMiddleware(uh.userDAO).GenerateJwtToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	refresh, err := middleware.NewAuthenticationMiddleware(uh.userDAO).GenerateJwtRefreshToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating refresh token"})
		return
	}

	body := userLoginResponse{
		Token:        token,
		RefreshToken: refresh,
	}

	c.JSON(http.StatusOK, body)
}
