package server

import (
	"eccom-mongo/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	docs "eccom-mongo/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	// Public routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/health", s.healthCtrl.HealthHandler)
	r.POST("/user", s.userCtrl.CreateUser)
	r.POST("/user/login", s.userCtrl.LoginUser)
	r.POST("/user/refresh-token", s.userCtrl.RefreshToken)
	r.GET("/product/{sku}", s.productCtrl.GetProduct)
	r.GET("/product/{terms}", s.productCtrl.SearchProduct)

	// Private routes
	r.Use(middleware.NewAuthenticationMiddleware(s.userDAO).Authenticate)

	r.POST("/address", s.addressCtrl.CreateAddress)
	r.GET("/address", s.addressCtrl.GetAllAddress)
	r.GET("/address/:id", s.addressCtrl.GetAddress)
	r.DELETE("/address/:id", s.addressCtrl.DeleteAddress)
	r.PUT("/address/:id/main", s.addressCtrl.MarkAddressAsMain)

	return r
}
