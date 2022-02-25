package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/routes"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {

	// Setup the API and all routes
	server := routes.SetupRoutes(logger)

	// Start the server with the configuration
	go func() {

		logger.Infoln("Server listening ... ")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// Context is used to inform the server it has 5 seconds to finish the request it is currently handling
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Do some actions here : close DB connections, ...
	logger.Infoln("SIGINT, SIGTERM received, will shut down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Infoln("Server forced to shutdown: ", err)
	} else {
		logger.Infoln("Application stopped gracefully !")
	}

}
