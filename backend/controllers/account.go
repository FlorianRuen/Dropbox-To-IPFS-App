package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AccountController interface {
	CheckLoginCredentials(c *gin.Context)
	GetUserDetails(c *gin.Context)
}

type accountController struct {
	logger       *logrus.Logger
	usersService services.UsersService
}

func NewAccountController(logger *logrus.Logger, usersService services.UsersService) AccountController {

	return accountController{
		logger:       logger,
		usersService: usersService,
	}
}

func (ctrl accountController) CheckLoginCredentials(c *gin.Context) {

	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	accountIdToCheck := &model.ValueToCheck{}

	if err := json.Unmarshal(bodyBytes, &accountIdToCheck); err != nil {
		panic(err)
	}

	// Get user by account id and return value
	accountFound, err := ctrl.usersService.CheckUserExistByAccountId(c, accountIdToCheck.Value)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// If no user found, return null
	if !accountFound {
		c.JSON(http.StatusNotFound, false)

	} else {
		c.JSON(http.StatusOK, true)
	}
}

func (ctrl accountController) GetUserDetails(c *gin.Context) {

	var bodyBytes []byte

	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	accountIdToCheck := &model.ValueToCheck{}

	if err := json.Unmarshal(bodyBytes, &accountIdToCheck); err != nil {
		panic(err)
	}

	// Get user by account id and return value
	user, err := ctrl.usersService.GetUserByAccountId(c, accountIdToCheck.Value)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// If no user found, return null
	if reflect.DeepEqual(user, model.User{}) == true {
		c.JSON(http.StatusNotFound, nil)

	} else {
		c.JSON(http.StatusOK, user)
	}
}
