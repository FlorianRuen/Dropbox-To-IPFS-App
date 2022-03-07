package model

import "encoding/json"

type EstuaryErrorResponse struct {
	Error string `json:"error"`
}

type EstuaryFileUploadResponse struct {
	CID       string      `json:"cid"`
	EstuaryId json.Number `json:"estuaryId"`
	Providers []string    `json:"providers"`
}
