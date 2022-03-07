package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
)

func ExtractEstuaryCallError(resp *http.Response) error {
	respBody, err := ioutil.ReadAll(resp.Body)
	var estuaryErrorResponse model.EstuaryErrorResponse

	if err != nil {
		return fmt.Errorf("Error reading response body: %s", err)
	}

	if err := json.Unmarshal(respBody, &estuaryErrorResponse); err != nil {
		return fmt.Errorf("Error reading response body: %s", err)
	}

	return fmt.Errorf("Error from Estuary: %s", estuaryErrorResponse.Error)
}
