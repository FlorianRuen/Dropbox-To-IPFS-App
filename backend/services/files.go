package services

import (
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/repository"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FilesService interface {
	TreatNewEvent(ctx *gin.Context, event *model.DropboxEvent)
}

type filesService struct {
	logger          *logrus.Logger
	usersRepository repository.UsersRepository
}

func NewFilesService(logger *logrus.Logger, usersRepository repository.UsersRepository) FilesService {

	return filesService{
		logger:          logger,
		usersRepository: usersRepository,
	}
}

func (s filesService) TreatNewEvent(ctx *gin.Context, event *model.DropboxEvent) {

	for _, account := range event.Address.Accounts {

		err, user := s.usersRepository.GetByAccountId(ctx, account)

		if err != nil {
			s.logger.Errorln(err)
		}

		utils.PrettyPrint(user)
	}
}
