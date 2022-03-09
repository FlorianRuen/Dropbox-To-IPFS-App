package model

import (
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type File struct {
	EstuaryId          json.Number    `json:"estuaryId" gorm:"column:estuary_id"`
	Name               string         `json:"name" gorm:"column:filename"`
	DropboxId          string         `json:"id" gorm:"column:dropbox_id"`
	Size               json.Number    `json:"size" gorm:"column:filesize"`
	DropboxContentHash string         `json:"content_hash" gorm:"column:dropbox_content_hash"`
	CID                string         `json:"cid" gorm:"column:cid"`
	Providers          pq.StringArray `json:"providers" gorm:"type:varchar(64)[]; column:providers"`
	UserAccountId      string         `json:"userAccountId" gorm:"column:user_account_id"`
	MigratedAt         time.Time      `json:"migratedAt" gorm:"column:migrated_at"`
}

func NewFile(dropboxFile DropboxFile, estuaryFile EstuaryFileUploadResponse, accountId string) *File {
	return &File{
		EstuaryId:          estuaryFile.EstuaryId,
		Name:               dropboxFile.Name,
		DropboxId:          dropboxFile.Id,
		Size:               dropboxFile.Size,
		DropboxContentHash: dropboxFile.ContentHash,
		CID:                estuaryFile.CID,
		Providers:          estuaryFile.Providers,
		UserAccountId:      accountId,
		MigratedAt:         time.Now(),
	}
}
