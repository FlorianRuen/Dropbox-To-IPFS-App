package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {

	logger.Infoln("Informations about the current run")

	// Setup the API and all routes
	// We will also setup the Auth Middleware inside (because we need services / repo to init it)
	// We will also setup all the cron tasks in the same time (because using services to get functions)
	server := route.SetupRoutes(config, db, filecoinApi, logger, Version, Buildtime)

}
