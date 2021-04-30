package api

import (
	"github.com/brianfromlife/golang-ecs/internal/api/handlers"
	"github.com/brianfromlife/golang-ecs/pkg/config"
	"github.com/brianfromlife/golang-ecs/pkg/data"
	"github.com/brianfromlife/golang-ecs/pkg/logger"
	"github.com/brianfromlife/golang-ecs/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	server *echo.Echo
	cfg    *config.Settings
}

func New(cfg *config.Settings, client *mongo.Client) *App {
	server := echo.New()

	server.Use(middleware.Recover())

	logger := logger.NewLogger(cfg)

	userProvider := data.NewUserProvider(cfg, client)
	userSvc := services.NewUserService(cfg, userProvider)

	healthHandler := handlers.NewHealthHandler()
	userHandler := handlers.NewUserHandler(logger, userSvc)

	server.GET("/v1/public/healthy", healthHandler.HealthCheck)
	server.POST("/v1/public/account/register", userHandler.Register)

	return &App{
		server: server,
		cfg:    cfg,
	}
}

func (a App) Start() {
	a.server.Start(":5000")
}
