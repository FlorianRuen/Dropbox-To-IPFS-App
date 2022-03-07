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

type DropboxFile struct {
	Name           string      `json:"name"`
	PathLower      string      `json:"path_lower"`
	PathDisplay    string      `json:"path_display"`
	Id             string      `json:"id"`
	ClientModified time.Time   `json:"client_modified"`
	ServerModified time.Time   `json:"server_modified"`
	Revision       string      `json:"rev"`
	Size           json.Number `json:"size"`
	IsDownloadable bool        `json:"is_downloadable"`
	ContentHash    string      `json:"content_hash"`
}
