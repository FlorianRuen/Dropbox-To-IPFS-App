package services

import (
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UsersService interface {
	InsertNewUser(c *gin.Context, user *model.DropboxUser) error
	GetUserByAccountId(c *gin.Context, accountId string) (error, model.DropboxUser)
}

type usersService struct {
	logger          *logrus.Logger
	usersRepository repository.UsersRepository
}

func NewUsersService(logger *logrus.Logger, usersRepository repository.UsersRepository) UsersService {

	return usersService{
		logger:          logger,
		usersRepository: usersRepository,
	}
}

func (s usersService) InsertNewUser(c *gin.Context, user *model.DropboxUser) error {
	return s.usersRepository.InsertNewUser(c, user)
}

func (s usersService) GetUserByAccountId(c *gin.Context, accountId string) (error, model.DropboxUser) {
	return s.usersRepository.GetByAccountId(c, accountId)
}
