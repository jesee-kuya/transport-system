package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jesee-kuya/transport-system/config"
	"github.com/jesee-kuya/transport-system/database"
	"github.com/jesee-kuya/transport-system/handler"
	"github.com/jesee-kuya/transport-system/middleware"
	"github.com/jesee-kuya/transport-system/repository"
	"github.com/jesee-kuya/transport-system/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)

	mw := middleware.NewMiddleware()

	t := handler.Transport{
		Middleware:  mw,
		AuthService: authService,
	}

	router := t.SetupRoutes()

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	// Start server in goroutine for graceful shutdown
	go func() {
		if err := router.Run(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

}
