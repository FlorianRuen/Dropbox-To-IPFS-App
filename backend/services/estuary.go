package services

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EstuaryService interface {
	SendFile(ctx *gin.Context, filename string) (model.EstuaryFileUploadResponse, error)
}

type estuaryService struct {
	config model.Config
	logger *logrus.Logger
}

func NewEstuaryService(config model.Config, logger *logrus.Logger) EstuaryService {

	return estuaryService{
		config: config,
		logger: logger,
	}
}

func (s estuaryService) SendFile(ctx *gin.Context, filename string) (model.EstuaryFileUploadResponse, error) {

	// Read the file from temp directory
	file, _ := os.Open((s.config.Files.TempFolder + filename))
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("data", filename)
	io.Copy(part, file)
	writer.Close()

	// Create a new request
	client := &http.Client{}

	req, err := http.NewRequest("POST", s.config.Estuary.UploadEndpoint, body)

	if err != nil {
		return model.EstuaryFileUploadResponse{}, err
	}

	req.Header.Set("Authorization", s.config.Estuary.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := client.Do(req)

	if err != nil {
		s.logger.Errorln(err)
		return model.EstuaryFileUploadResponse{}, err
	}

	// If response status code isn't 200, threat it as an error
	// Otherwise get uploaded file metadata
	if response.StatusCode != 200 {

		err := utils.ExtractEstuaryCallError(response)

		s.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Error while sending file to Estuary API with message")

		return model.EstuaryFileUploadResponse{}, err

	} else {

		respBody, err := ioutil.ReadAll(response.Body)
		estuaryUploadedFile := &model.EstuaryFileUploadResponse{}

		if err != nil {

			s.logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("Error while sending file to Estuary API with message")

			return model.EstuaryFileUploadResponse{}, err
		}

		// Unmarshal the response content to get the error message
		if err := json.Unmarshal(respBody, &estuaryUploadedFile); err != nil {

			s.logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("Error while sending file to Estuary API with message")

			return model.EstuaryFileUploadResponse{}, err
		}

		return *estuaryUploadedFile, nil

	}
}
