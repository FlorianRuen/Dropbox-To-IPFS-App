package utils

import (
	"encoding/json"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

func ConvertMetadataToFile(fileMetadata files.IsMetadata) (file model.DropboxFile, err error) {

	b, err := json.MarshalIndent(fileMetadata, "", "  ")

	if err != nil {
		return file, err
	}

	if err := json.Unmarshal(b, &file); err != nil {
		return file, err
	}

	return file, nil
}
