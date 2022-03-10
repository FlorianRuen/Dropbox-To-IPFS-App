package repository

import (
	"time"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UsersRepository interface {
	UpdateOrInsertUser(c *gin.Context, user *model.User) error
	GetByAccountId(c *gin.Context, accountId string) (model.User, error)
	CheckUserExistByAccountId(ctx *gin.Context, accountId string) (bool, error)
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

func (r usersRepository) UpdateOrInsertUser(c *gin.Context, user *model.User) error {

	request := r.dbClient.Model(&user).
		Table("users").
		Where("account_id = ?", user.AccountId).
		Updates(map[string]interface{}{
			"uid":          user.UniqueIdentifer,
			"scopes":       user.Scopes,
			"expires_in":   user.ExpiresIn,
			"access_token": user.AccessToken,
			"retrieved_at": time.Now(),
			"firstname":    user.Firstname,
			"lastname":     user.Lastname,
			"email":        user.Email,
		})

	// If no rows updated = user doesn't exist, so create it in database
	if request.RowsAffected == 0 {

		request = r.dbClient.Scopes().
			Table("users").
			Create(user)

		return request.Error

	} else {
		return request.Error
	}
}

func (r usersRepository) GetByAccountId(ctx *gin.Context, accountId string) (model.User, error) {

	var user model.User

	req := r.dbClient.Scopes().
		Table("users").
		Where("account_id = ?", accountId).
		Find(&user)

	return user, req.Error
}

func (r usersRepository) CheckUserExistByAccountId(ctx *gin.Context, accountId string) (bool, error) {

	var exists int64

	req := r.dbClient.Scopes().
		Table("users").
		Where("account_id = ?", accountId).
		Count(&exists)

	return exists > 0, req.Error
}

func (r usersRepository) GetCursor(ctx *gin.Context, accountId string) (model.Cursor, error) {

	var cursor model.Cursor

	req := r.dbClient.Scopes().
		Table("cursors").
		Where("account_id = ?", accountId).
		Find(&cursor)

	return cursor, req.Error
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
