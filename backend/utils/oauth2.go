package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/sirupsen/logrus"
)

func GetAccessToken(ctx context.Context, logger *logrus.Logger, code string) (*model.DropboxUser, error) {

	data := url.Values{
		"client_id":     {"c3hbpngaqu240bf"},
		"client_secret": {"jwqvwouc0e6daeu"},
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {"http://localhost:3200/api/dropbox/oauth_callback"},
	}

	response, err := http.PostForm("https://api.dropboxapi.com/oauth2/token", data)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	callback_response := &model.DropboxUser{}

	if err != nil {

		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Error while getting challenge from Estuary Api with message")

		return callback_response, err
	}

	// Unmarshal the response content to get the error message
	if err := json.Unmarshal(body, &callback_response); err != nil {

		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Error while unmarshal the response from Estuary while checking challenge")

		return callback_response, err
	}

	// Set current time in the struct
	callback_response.RetrievedAt = time.Now()
	return callback_response, nil
}
