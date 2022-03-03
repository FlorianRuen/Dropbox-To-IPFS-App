package model

import (
	"encoding/json"
	"time"
)

type DropboxEvent struct {
	Delta   DropboxDelta       `json:"delta"`
	Address DropboxFoldersList `json:"list_folder"`
}

type DropboxFoldersList struct {
	Accounts []string `json:"accounts"`
}

type DropboxDelta struct {
	Users []json.Number `json:"users"`
}

type DropboxUser struct {
	AccessToken     string      `json:"access_token" gorm:"column:access_token"`
	TokenType       string      `json:"token_type" gorm:"column:token_type"`
	RetrievedAt     time.Time   `json:"retrieved_at" gorm:"column:retrieved_at"`
	ExpiresIn       json.Number `json:"expires_in" gorm:"column:expires_in"`
	Scopes          string      `json:"scope" gorm:"column:scopes"`
	UniqueIdentifer string      `json:"uid" gorm:"column:uid"`
	AccountId       string      `json:"account_id" gorm:"column:account_id"`
}
