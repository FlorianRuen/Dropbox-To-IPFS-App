package repository

import (
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UsersRepository interface {
	InsertNewUser(c *gin.Context, user *model.DropboxUser) error
	GetByAccountId(c *gin.Context, accountId string) (error, model.DropboxUser)
}

type usersRepository struct {
	logger   *logrus.Logger
	dbClient *gorm.DB
}

func NewUsersRepository(logger *logrus.Logger, dbClient *gorm.DB) UsersRepository {

	return usersRepository{
		logger:   logger,
		dbClient: dbClient,
	}
}

func (r usersRepository) InsertNewUser(c *gin.Context, user *model.DropboxUser) error {

	request := r.dbClient.Scopes().
		Table("users").
		Create(&user)

	if request.Error != nil {
		return request.Error
	}

	return nil
}

func (r usersRepository) GetByAccountId(ctx *gin.Context, accountId string) (error, model.DropboxUser) {

	var user model.DropboxUser

	// By default, we select all clients (including ones who are archieved)
	req := r.dbClient.Scopes().
		Table("users").
		Where("account_id = ?", accountId)

	request := req.Find(&user)
	return request.Error, user
}
