package controller

import (
	"eccom-mongo/internal/database"
	"eccom-mongo/internal/middleware"
	"eccom-mongo/internal/models"
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
	Nome            string `json:"nome" binding:"required"`
	Sobrenome       string `json:"sobrenome" binding:"required"`
	CPF             string `json:"documento" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type userLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "password and confirm password do not match"})
		return
	}

	passwordHash, err := hashBcrypt(userCreate.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing password"})
		return
	}

	user, err := uh.userDAO.FindByEmail(c, userCreate.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking if user exists"})
		return
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	user, err = uh.userDAO.FindByCPF(c, userCreate.CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking if CPF exists"})
		return
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CPF already exists"})
		return
	}

	newUser := models.User{
		Nome:         userCreate.Nome,
		CPF:          userCreate.CPF,
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

func hashBcrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// @Summary Login user
// @Description Login user
// @Tags user
// @Accept json
// @Produce json
// @Param user body userLoginRequest true "User data"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /user/login [post]
func (uh *userHandler) LoginUser(c *gin.Context) {
	var userLogin userLoginRequest
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userDAO.FindByEmail(c, userLogin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking if user exists"})
		return
	}

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userLogin.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
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

	body := gin.H{
		"token":         token,
		"refresh_token": refresh,
	}

	c.JSON(http.StatusOK, body)
}
