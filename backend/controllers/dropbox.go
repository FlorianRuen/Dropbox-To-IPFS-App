package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/services"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/alexellis/hmac"
)

type DropboxController interface {
	ValidDropboxWebsocketChallenge(c *gin.Context)
	AuthentificationCallback(c *gin.Context)
	ReceivedDropboxEventsNotifications(c *gin.Context)
}

type dropboxController struct {
	logger         *logrus.Logger
	usersService   services.UsersService
	filesService   services.FilesService
	dropboxService services.DropboxService
}

func NewDropboxController(logger *logrus.Logger, usersService services.UsersService, filesService services.FilesService,
	dropboxService services.DropboxService) DropboxController {

	return dropboxController{
		logger:         logger,
		usersService:   usersService,
		filesService:   filesService,
		dropboxService: dropboxService,
	}
}

func (ctrl dropboxController) ValidDropboxWebsocketChallenge(c *gin.Context) {
	challenge := c.Request.URL.Query().Get("challenge")
	c.Header("Content-Type", "text/plain")
	c.Header("X-Content-Type-Options", "nosniff")
	c.String(http.StatusOK, challenge)
}

func (ctrl dropboxController) AuthentificationCallback(c *gin.Context) {

	// Get the token using the code in the authorize response
	code := c.Request.FormValue("code")
	callback_response, err := utils.GetAccessToken(c, ctrl.logger, code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// Get the user's account details by calling Dropbox API
	userAccountDetails, err := ctrl.dropboxService.GetUserAccount(c, callback_response.AccessToken, callback_response.AccountId)
	callback_response.Details = userAccountDetails

	// Store the access token in Redis database
	err = ctrl.usersService.InsertNewUser(c, callback_response)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	http.Redirect(c.Writer, c.Request, "http://www.golang.org", 301)
}

func (ctrl dropboxController) ReceivedDropboxEventsNotifications(c *gin.Context) {

	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	dropboxEvent := &model.DropboxEvent{}

	if err := json.Unmarshal(bodyBytes, &dropboxEvent); err != nil {
		panic(err)
	}

	// Make sure this is a valid request from Dropbox
	signature := c.Request.Header.Get("X-Dropbox-Signature")
	valid := hmac.Validate(bodyBytes, signature, "jwqvwouc0e6daeu")

	if valid != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid signature"})
	}

	// Parse data and threat in another thread
	ctrl.filesService.TreatNewEvent(c, dropboxEvent)

	// Because we got time to return the response
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
