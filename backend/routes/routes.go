package routes

import (
	"net/http"
	"time"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/controllers"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/repository"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CustomLogger(logger *logrus.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {

		// Get latency time for the request to return a result
		startTime := time.Now()
		c.Next()
		endTime := time.Now()

		// Set the log level for GIN logs
		logger.SetLevel(logrus.TraceLevel)

		logger.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"latency":    endTime.Sub(startTime),
			"endpoint":   c.Request.RequestURI,
			"statusCode": c.Writer.Status(),
			"realIP":     c.Request.Header.Get("X-Real-IP"),
		}).Info("GIN Request on API with parameters")
	}
}

func SetupRoutes(logger *logrus.Logger, dBClient *gorm.DB) http.Server {

	logger.Info("Setup all REST API endpoints ...")

	// Define all the API endpoints
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)

	server := &http.Server{
		Addr:    ":3200",
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

	// Init : repository for the DB exchanges
	usersRepository := repository.NewUsersRepository(logger, dBClient)

	// Init : services for the logic
	usersService := services.NewUsersService(logger, usersRepository)
	filesService := services.NewFilesService(logger, usersRepository)

	// Init : controllers to handle external requests
	dropboxController := controllers.NewDropboxController(logger, usersService, filesService)

	// Configure logger format for GIN
	router.Use(CustomLogger(logger))

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))

	// Setup all API routes
	apiV1 := router.Group("/api/dropbox")
	{
		// Route to install the Dropbox app
		apiV1.GET("/events", dropboxController.ValidDropboxWebsocketChallenge)
		apiV1.GET("/oauth_callback", dropboxController.AuthentificationCallback)

		// Route to receive all events notifications
		apiV1.POST("/events", dropboxController.ReceivedDropboxEventsNotifications)
	}

	return *server
}
