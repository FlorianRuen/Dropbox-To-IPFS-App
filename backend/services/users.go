package services

import (
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UsersService interface {
	UpdateOrInsertUser(c *gin.Context, user *model.User) error
	GetUserByAccountId(c *gin.Context, accountId string) (model.User, error)
	CheckUserExistByAccountId(c *gin.Context, accountId string) (bool, error)
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

func (s usersService) UpdateOrInsertUser(c *gin.Context, user *model.User) error {
	return s.usersRepository.UpdateOrInsertUser(c, user)
}

func (s usersService) GetUserByAccountId(c *gin.Context, accountId string) (model.User, error) {
	return s.usersRepository.GetByAccountId(c, accountId)
}

func (s usersService) CheckUserExistByAccountId(c *gin.Context, accountId string) (bool, error) {
	return s.usersRepository.CheckUserExistByAccountId(c, accountId)
}
