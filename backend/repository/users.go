package repository

import (
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UsersRepository interface {
	InsertNewUser(c *gin.Context, user *model.User) error
	GetByAccountId(c *gin.Context, accountId string) (model.User, error)
	GetCursor(ctx *gin.Context, accountId string) (model.Cursor, error)
	StoreCursor(ctx *gin.Context, accountId string, cursor string) error
	UpdateCursor(ctx *gin.Context, accountId string, cursor string) error
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

func (r usersRepository) InsertNewUser(c *gin.Context, user *model.User) error {

	request := r.dbClient.Scopes().
		Table("users").
		Create(&user)

	if request.Error != nil {
		return request.Error
	}

	return nil
}

func (r usersRepository) GetByAccountId(ctx *gin.Context, accountId string) (model.User, error) {

	var user model.User

	req := r.dbClient.Scopes().
		Table("users").
		Where("account_id = ?", accountId)

	request := req.Find(&user)
	return user, request.Error
}

func (r usersRepository) GetCursor(ctx *gin.Context, accountId string) (model.Cursor, error) {

	var cursor model.Cursor

	req := r.dbClient.Scopes().
		Table("cursors").
		Where("account_id = ?", accountId)

	request := req.Find(&cursor)
	return cursor, request.Error
}

func (r usersRepository) StoreCursor(ctx *gin.Context, accountId string, cursor string) error {

	request := r.dbClient.Scopes().
		Table("cursors").
		Create(model.Cursor{
			AccountId: accountId,
			Cursor:    cursor,
		})

	if request.Error != nil {
		return request.Error
	}

	return nil
}

func (r usersRepository) UpdateCursor(ctx *gin.Context, accountId string, cursor string) error {

	request := r.dbClient.Scopes().
		Table("cursors").
		Where("account_id = ?", accountId).
		Updates(model.Cursor{
			Cursor: cursor,
		})

	if request.Error != nil {
		return request.Error
	}

	return nil
}
