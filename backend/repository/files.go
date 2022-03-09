package repository

import (
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FilesRepository interface {
	InsertNewFile(c *gin.Context, file *model.File) error
	GetFilesForCurrentUser(ctx *gin.Context, accountId string) ([]model.File, error)
}

type filesRepository struct {
	logger   *logrus.Logger
	dbClient *gorm.DB
}

func NewFilesRepository(logger *logrus.Logger, dbClient *gorm.DB) FilesRepository {

	return filesRepository{
		logger:   logger,
		dbClient: dbClient,
	}
}

func (r filesRepository) InsertNewFile(c *gin.Context, file *model.File) error {

	request := r.dbClient.Scopes().
		Table("migrated_files").
		Create(&file)

	if request.Error != nil {
		return request.Error
	}

	return nil
}

func (r filesRepository) GetFilesForCurrentUser(ctx *gin.Context, accountId string) ([]model.File, error) {

	var files []model.File

	req := r.dbClient.Scopes().
		Table("migrated_files").
		Where("user_account_id = ?", accountId).
		Find(&files)

	return files, req.Error
}
