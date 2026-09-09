package main

import (
	"fmt"
	"go-landing-page/internal/api"
	"go-landing-page/internal/api/middleware"
	"go-landing-page/internal/database"
	"go-landing-page/internal/repositories"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"
)

type Server struct {
	port string

	db database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	userRepo := repositories.NewUserRepository(s.db)
	userHandler := api.NewUserHandler(*userRepo)
	eventRepo := repositories.NewEventRepository(s.db)
	eventHandler := api.NewEventAPI(eventRepo)
	leadRepo := repositories.NewLeadRepository(s.db)
	leadHandler := api.NewLeadApi(leadRepo)
	pageViewRepo := repositories.NewPageViewRepository(s.db)
	pageViewHandler := api.NewPageViewApi(pageViewRepo)
	metricsRepo := repositories.NewMetricsRepository(s.db)
	metricsHandler := api.NewMetricsAPI(metricsRepo)

	apiRoutes := e.Group("/api")
	apiRoutes.POST("/hash-password", userHandler.HashPassword)

	authRoutes := apiRoutes.Group("/auth")
	authRoutes.POST("/login", userHandler.Login)

	apiRoutes.POST("/events", eventHandler.CreateEvent)
	apiRoutes.POST("/leads", leadHandler.CreateLead)
	apiRoutes.POST("/page-views", pageViewHandler.CreatePageView)

	protected := apiRoutes.Group("", middleware.JWTMiddleware)
	protected.GET("/metrics/summary", metricsHandler.GetSummary)
	protected.GET("/events", eventHandler.ListEvents)
	protected.GET("/leads", leadHandler.ListLeads)
	protected.GET("/page-views", pageViewHandler.ListPageViews)

	return e
}

func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	s := &Server{
		port: fmt.Sprintf(":%d", port),
		db:   database.NewConnectionDatabase(),
	}

	server := &http.Server{
		Addr:         s.port,
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	fmt.Println("Server running on port", s.port)

	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
