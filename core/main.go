package main

import (
	"go-basic/application/util"
	"go-basic/infrastructure/db"
	"go-basic/presentation/route"
	"go-basic/security"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDatabase()

	r := gin.Default()

	r.Use(cors.Default())

	exemptRoutes := []string{
		"/api/health",
		"/api/auth/login",
		"/api/auth/register",
	}

	logger := *util.NewLogger("GUARD")
	guardMiddleware := security.Guard(exemptRoutes, logger)
	if guardMiddleware == nil {
		log.Fatal("Security middleware initialization failed")
	}
	r.Use(guardMiddleware)

	route.RegisterIndexRoutes(r)
	route.RegisterAuthRoutes(r)

	log.Println("Server is running on port 3000")
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
