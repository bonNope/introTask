package app

import (
	"github.com/bonNope/introTask/internal/app/handler"
	"github.com/bonNope/introTask/internal/app/middleware"
	"github.com/bonNope/introTask/internal/app/service"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	h *handler.Handler
	s *service.Service
	g *gin.Engine
}

func New() *App {
	a := &App{}

	a.s = service.NewService()
	a.h = handler.NewHandler(a.s)
	a.g = gin.Default()

	a.g.GET("/when/:year", middleware.CheckHeader(), a.h.CountDays)

	return a
}

func (a App) Run() error {
	err := a.g.Run()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
