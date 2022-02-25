package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type DropboxController interface {
	ValidDropboxWebsocketChallenge(c *gin.Context)
	ReceivedDropboxEventsNotifications(c *gin.Context)
}

type dropboxController struct {
	logger *logrus.Logger
}

func NewDropboxController(logger *logrus.Logger) DropboxController {

	return dropboxController{
		logger: logger,
	}
}

func (ctrl dropboxController) ValidDropboxWebsocketChallenge(c *gin.Context) {
	challenge := c.Request.URL.Query().Get("challenge")
	c.Header("Content-Type", "text/plain")
	c.Header("X-Content-Type-Options", "nosniff")
	c.String(http.StatusOK, challenge)
}

func (ctrl dropboxController) ReceivedDropboxEventsNotifications(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	utils.PrettyPrint(jsonData)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
