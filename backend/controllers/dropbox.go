package controllers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/services"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	url_to_redirect := "http://localhost:4041/callback/" + callback_response.AccountId
	http.Redirect(c.Writer, c.Request, url_to_redirect, 301)
}

func (ctrl dropboxController) ReceivedDropboxEventsNotifications(c *gin.Context) {

	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	dropboxEvent := &model.DropboxEvent{}

	if err := json.Unmarshal(bodyBytes, &dropboxEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode event informations"})
		return
	}

	// Make sure this is a valid request from Dropbox
	hexSignature, err := hex.DecodeString(c.Request.Header.Get("X-Dropbox-Signature"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode signature header"})
	}

	// Encode the secret token to compare to signature
	expectedSignature := hmac.New(sha256.New, []byte("jwqvwouc0e6daeu")).Sum(nil)
	valid := hmac.Equal(hexSignature, expectedSignature)

	if valid {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid signature"})
		return
	}

	// For each account, start a new thread to process the event
	// Because we have a limited time to return a valid response to Dropbox API
	for _, account := range dropboxEvent.Address.Accounts {

		go func() {
			ctrl.logger.Infoln("Launch new goroutine to process event for account", account)
			ctrl.filesService.ProcessUserEvent(c, account)
		}()

	}

	// Because we got time to return the response
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
