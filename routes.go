package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func SetupRoutes(logger *logrus.Logger) http.Server {

	logger.Info("Setup all REST API endpoints ...")

	// Define all the API endpoints
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	server := &http.Server{
		Addr:    ":" + config.API.ListenPort,
		Handler: router,
	}

	router.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, PATCH"},
			AllowHeaders:     []string{"Content-Type, Upgrade, Content-Length, Accept-Encoding, Host, Authorization, accept, Origin, Cache-Control, X-Requested-With"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)

	// Configure logger format for GIN
	router.Use(CustomLogger(config, logger))

	// Setup all API routes
	apiV1 := router.Group("/api")
	{
		apiV1.GET("/version", appController.GetVersionAndBuildTime)
	}

}
