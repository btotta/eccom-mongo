package dtos

import "eccom-mongo/internal/models"

type UserDTO struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}

type UserRegisterDTO struct {
	Name            string `json:"name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Document        string `json:"document" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponseDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserDTO(user *models.User) *UserDTO {
	return &UserDTO{
		ID:       user.ID.Hex(),
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
}
