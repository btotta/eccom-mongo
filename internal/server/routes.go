package server

import (
	"eccom-mongo/internal/controller"
	"eccom-mongo/internal/database"
	"eccom-mongo/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	docs "eccom-mongo/cmd/api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	userController := controller.NewUserController(database.NewUserDAO(s.db.GetDB()))
	healthController := controller.NewHealthController(&s.db)

	// Public routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/health", healthController.HealthHandler)
	r.POST("/user", userController.CreateUser)
	r.POST("/user/login", userController.LoginUser)

	// Private routes
	r.Use(middleware.NewAuthenticationMiddleware(database.NewUserDAO(s.db.GetDB())).Authenticate)

	return r
}
