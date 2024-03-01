package controller

import (
	"eccom-mongo/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	db database.Service
}

func NewHealthController(db *database.Service) *HealthHandler {
	return &HealthHandler{db: *db}
}

// @Summary Health check
// @Description Check if the server is healthy
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /health [get]
func (h *HealthHandler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, h.db.Health())
}
