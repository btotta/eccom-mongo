package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"eccom-mongo/internal/controller"
	"eccom-mongo/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port        int
	userDAO     database.UserDAOInterface
	addressDAO  database.AddressDAOInterface
	db          database.Service
	healthCtrl  controller.HealthHandler
	userCtrl    controller.UserHandler
	addressCtrl controller.AddressHandler
}

func NewServer() *http.Server {

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Erro ao converter a porta:", err)
		os.Exit(1)
	}

	db := database.New()

	// Initialize DAOs
	userDAO := database.NewUserDAO(db.GetDB())
	addressDAO := database.NewAddressDAO(db.GetDB())

	// Initialize controllers
	healthCtrl := controller.NewHealthController(db)
	userCtrl := controller.NewUserController(userDAO)
	addressCtrl := controller.NewAddressController(addressDAO, userDAO)

	server := &Server{
		port:        port,
		db:          db,
		userDAO:     userDAO,
		addressDAO:  addressDAO,
		healthCtrl:  healthCtrl,
		userCtrl:    userCtrl,
		addressCtrl: addressCtrl,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%d", server.port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
