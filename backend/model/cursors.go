package model

type Cursor struct {
	AccountId string `json:"delta" gorm:"column:account_id"`
	Cursor    string `json:"list_folder" gorm:"column:cursor"`
}
