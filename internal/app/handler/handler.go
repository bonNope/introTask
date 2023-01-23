package handler

import (
	"github.com/bonNope/introTask/internal/app/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CountDays(c *gin.Context) {
	yearParam := c.Param("year")

	year, err := strconv.Atoi(yearParam)
	if err != nil {
		log.Print("invalid input year")
		return
	}

	d := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	result := h.service.CalculateDaysGoneOrLeft(d)

	c.String(http.StatusOK, result)
}
