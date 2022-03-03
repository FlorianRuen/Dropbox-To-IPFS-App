package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/database"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/routes"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {

	// Get executable path
	dir, err_dir := filepath.Abs(filepath.Dir(os.Args[0]))

	if err_dir != nil {

		logger.WithFields(logrus.Fields{
			"err":      err_dir,
			"location": dir,
		}).Panicln("Can't find executable path with message with message")

		panic(err_dir)
	}

	// Check if config file found in both location
	configFilePath := dir + "/config.toml"

	if _, err := os.Stat(dir + "/config.toml"); errors.Is(err, os.ErrNotExist) {

		if _, err := os.Stat("config.toml"); errors.Is(err, os.ErrNotExist) {

			logger.WithFields(logrus.Fields{
				"err":      err,
				"location": dir,
			}).Panicln("Couldn't open config.toml file with message")

		} else {
			configFilePath = "config.toml"
		}

	}

	// First we need to load the config.toml
	var config model.Config

	meta, err := toml.DecodeFile(configFilePath, &config)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"err":      err,
			"location": dir,
		}).Panicln("Couldn't open config.toml file with message")
	}

	// Check all the keys from config toml has been decoded
	errConfigFile := model.CheckValidConfig(config, meta)

	if errConfigFile != nil {

		logger.WithFields(logrus.Fields{
			"err": errConfigFile,
		}).Errorln("Invalid config file")

		logger.Fatalf("Process aborted")
	}

	// Database : init settings and test connexion
	dbClient := database.ConnectToDatabase(config.Database)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Unable to connect to Redis local server, check if it's reachable")

		logger.Fatalf("Process aborted")
	}

	// Setup the API and all routes
	server := routes.SetupRoutes(logger, dbClient)

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
