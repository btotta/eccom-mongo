package middleware

import (
	"eccom-mongo/internal/database"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	secretKey = []byte(os.Getenv("JWT_SECRET"))
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type AuthenticationMiddleware interface {
	Authenticate(c *gin.Context)
	GenerateJwtToken(email string) (string, error)
	GenerateJwtRefreshToken(email string) (string, error)
}

type authenticationMiddleware struct {
	userDAO database.UserDAOInterface
}

func NewAuthenticationMiddleware(userDAO database.UserDAOInterface) AuthenticationMiddleware {
	return &authenticationMiddleware{
		userDAO: userDAO,
	}
}

func (am *authenticationMiddleware) Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação não fornecido"})
		c.Abort()
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Falha ao analisar o token de autenticação"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação inválido"})
		c.Abort()
		return
	}

	user, err := am.userDAO.FindByEmail(c, claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário no banco de dados"})
		c.Abort()
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}

func (am *authenticationMiddleware) GenerateJwtToken(email string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (am *authenticationMiddleware) GenerateJwtRefreshToken(email string) (string, error) {
	expirationTime := time.Now().Add(6 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
