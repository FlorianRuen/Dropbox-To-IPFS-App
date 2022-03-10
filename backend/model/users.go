package model

import (
	"encoding/json"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/users"
)

type User struct {
	Firstname       string      `json:"firstName" gorm:"column:firstname"`
	Lastname        string      `json:"lastName" gorm:"column:lastname"`
	Email           string      `json:"email" gorm:"column:email"`
	AccessToken     string      `json:"access_token" gorm:"column:access_token"`
	TokenType       string      `json:"token_type" gorm:"column:token_type"`
	RetrievedAt     time.Time   `json:"retrieved_at" gorm:"column:retrieved_at"`
	ExpiresIn       json.Number `json:"expires_in" gorm:"column:expires_in"`
	Scopes          string      `json:"scope" gorm:"column:scopes"`
	UniqueIdentifer string      `json:"uid" gorm:"column:uid"`
	AccountId       string      `json:"account_id" gorm:"column:account_id"`
}

func NewUser(dropboxUser *users.BasicAccount, callbackUser *User) *User {
	return &User{
		AccessToken:     callbackUser.AccessToken,
		TokenType:       callbackUser.TokenType,
		RetrievedAt:     callbackUser.RetrievedAt,
		ExpiresIn:       callbackUser.ExpiresIn,
		Scopes:          callbackUser.Scopes,
		UniqueIdentifer: callbackUser.UniqueIdentifer,
		AccountId:       callbackUser.AccountId,
		Firstname:       dropboxUser.Name.GivenName,
		Lastname:        dropboxUser.Name.Surname,
		Email:           dropboxUser.Email,
	}
}
