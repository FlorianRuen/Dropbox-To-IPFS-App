package services

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

type DropboxService interface {
	GetUserAccount(ctx *gin.Context, token string, accountId string) (*users.BasicAccount, error)
	GetFiles(ctx *gin.Context, token string) (*files.ListFolderResult, error)
	GetFilesContinue(ctx *gin.Context, token string, cursor string) (*files.ListFolderResult, error)
	DownloadFile(ctx *gin.Context, token string, path string) error
}

type dropboxService struct {
	logger          *logrus.Logger
	dropboxLogLevel dropbox.LogLevel
}

func NewDropboxService(logger *logrus.Logger) DropboxService {

	return dropboxService{
		logger:          logger,
		dropboxLogLevel: dropbox.LogDebug,
	}
}

func (s dropboxService) GetUserAccount(ctx *gin.Context, token string, accountId string) (*users.BasicAccount, error) {

	config := dropbox.Config{
		Token:    token,
		LogLevel: s.dropboxLogLevel,
	}

	arg := users.NewGetAccountArg(accountId)
	return users.New(config).GetAccount(arg)
}

func (s dropboxService) GetFiles(ctx *gin.Context, token string) (*files.ListFolderResult, error) {

	config := dropbox.Config{
		Token:    token,
		LogLevel: s.dropboxLogLevel,
	}

	arg := files.NewListFolderArg("")
	return files.New(config).ListFolder(arg)
}

func (s dropboxService) GetFilesContinue(ctx *gin.Context, token string, cursor string) (*files.ListFolderResult, error) {

	config := dropbox.Config{
		Token:    token,
		LogLevel: s.dropboxLogLevel,
	}

	arg := files.NewListFolderContinueArg(cursor)
	return files.New(config).ListFolderContinue(arg)
}

func (s dropboxService) DownloadFile(ctx *gin.Context, token string, path string) error {

	config := dropbox.Config{
		Token:    token,
		LogLevel: s.dropboxLogLevel,
	}

	metadata, content, err := files.New(config).Download(
		files.NewDownloadArg(path),
	)

	fo, err := os.Create("temp_content/" + metadata.Name)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)

	for {
		n, err := content.Read(buf)

		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	return nil
}
