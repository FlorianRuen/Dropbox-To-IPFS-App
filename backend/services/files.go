package services

import (
	"reflect"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/repository"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/utils"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FilesService interface {
	ProcessUserEvent(ctx *gin.Context, account string)
	GetFilesForCurrentUser(ctx *gin.Context, account string) ([]model.File, error)
}

type filesService struct {
	logger          *logrus.Logger
	dropboxService  DropboxService
	estuaryService  EstuaryService
	filesRepository repository.FilesRepository
	usersRepository repository.UsersRepository
}

func NewFilesService(logger *logrus.Logger, dropboxService DropboxService, estuaryService EstuaryService,
	filesRepository repository.FilesRepository, usersRepository repository.UsersRepository) FilesService {

	return filesService{
		logger:          logger,
		dropboxService:  dropboxService,
		estuaryService:  estuaryService,
		filesRepository: filesRepository,
		usersRepository: usersRepository,
	}
}

func (s filesService) ProcessUserEvent(ctx *gin.Context, account string) {

	// Retrieve token by account id in event request
	user, err := s.usersRepository.GetByAccountId(ctx, account)

	if err != nil {
		s.logger.Errorln(err)
		return
	}

	// Get cursor stored in database for user concerned by event
	// This will allow us to get only new files and ignore old ones
	cursor, err := s.usersRepository.GetCursor(ctx, account)

	if err != nil {
		s.logger.Errorln(err)
		return
	}

	// Depending if cursor exist, browse the dropbox app folder for the first time or continue
	var filesList *files.ListFolderResult

	if reflect.DeepEqual(cursor, model.Cursor{}) == true {

		filesList, err = s.dropboxService.GetFiles(ctx, user.AccessToken)

		if err != nil {
			s.logger.Errorln(err)
			return
		}

	} else {

		filesList, err = s.dropboxService.GetFilesContinue(ctx, user.AccessToken, cursor.Cursor)

		if err != nil {
			s.logger.Errorln(err)
			return
		}
	}

	// For each file, download and send it to IPFS using Estuary API
	// In the POC version, use only files, ignore folders
	for _, fileMetadata := range filesList.Entries {

		file, err := utils.ConvertMetadataToFile(fileMetadata)

		if err != nil {
			s.logger.Errorln(err)
			return
		}

		// Download the file from Dropbox
		s.logger.Infoln("Downloading file from Dropbox ...")
		err = s.dropboxService.DownloadFile(ctx, user.AccessToken, file.PathLower)

		if err != nil {
			s.logger.Errorln(err)
			return
		}

		// Send over Estuary HTTP API
		s.logger.Infoln("Send file to IPFS using Estuary ...")
		uploadedFile, err := s.estuaryService.SendFile(ctx, file.Name)

		if err != nil {
			s.logger.Errorln(err)
			return
		}

		// Create a file with all uploaded informations and save to database
		filecoinFile := model.NewFile(file, uploadedFile, user.AccountId)
		err = s.filesRepository.InsertNewFile(ctx, filecoinFile)

		if err != nil {
			s.logger.Errorln(err)
			return
		}

		s.logger.Infoln("File migrated from Dropbox to IPFS !")
	}

	// To avoid browse all the dropbox folder again, store the cursor in database
	// Or update the existing cursor to get only new files on the next attempt
	if reflect.DeepEqual(cursor, model.Cursor{}) == true {
		err = s.usersRepository.StoreCursor(ctx, account, filesList.Cursor)

	} else {
		err = s.usersRepository.UpdateCursor(ctx, account, filesList.Cursor)
	}
}

func (s filesService) GetFilesForCurrentUser(ctx *gin.Context, account string) ([]model.File, error) {

	files, err := s.filesRepository.GetFilesForCurrentUser(ctx, account)

	if err != nil {

		s.logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Error while getting all migrated files for current user")

		return []model.File{}, err
	}

	return files, nil
}
