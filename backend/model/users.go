package model

import (
	"encoding/json"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

type User struct {
	Details         *users.BasicAccount `json:"account_details" gorm:"-"`
	AccessToken     string              `json:"access_token" gorm:"column:access_token"`
	TokenType       string              `json:"token_type" gorm:"column:token_type"`
	RetrievedAt     time.Time           `json:"retrieved_at" gorm:"column:retrieved_at"`
	ExpiresIn       json.Number         `json:"expires_in" gorm:"column:expires_in"`
	Scopes          string              `json:"scope" gorm:"column:scopes"`
	UniqueIdentifer string              `json:"uid" gorm:"column:uid"`
	AccountId       string              `json:"account_id" gorm:"column:account_id"`
}
