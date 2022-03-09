package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FilesController interface {
	GetFilesForCurrentUser(c *gin.Context)
}

type filesController struct {
	logger       *logrus.Logger
	usersService services.UsersService
	filesService services.FilesService
}

func NewFilesController(logger *logrus.Logger, usersService services.UsersService, filesService services.FilesService) FilesController {

	return filesController{
		logger:       logger,
		usersService: usersService,
		filesService: filesService,
	}
}

func (ctrl filesController) GetFilesForCurrentUser(c *gin.Context) {

	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	accountIdToCheck := &model.ValueToCheck{}

	if err := json.Unmarshal(bodyBytes, &accountIdToCheck); err != nil {
		panic(err)
	}

	// Get files by account id and return value
	files, err := ctrl.filesService.GetFilesForCurrentUser(c, accountIdToCheck.Value)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, files)
}
