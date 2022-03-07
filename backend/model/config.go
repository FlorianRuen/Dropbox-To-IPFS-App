package model

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	API      APIConfig      `json:"apiConfig"`
	Database DatabaseConfig `json:"databaseConfig"`
	Files    FilesConfig    `json:"filesConfig"`
	Estuary  EstuaryConfig  `json:"estuaryConfig"`
}

type APIConfig struct {
	ListenPort string `json:"listenPort"`
}

type DatabaseConfig struct {
	User     string `json:"databaseUser"`
	Password string `json:"databasePassword"`
	Host     string `json:"databaseHost"`
	Port     string `json:"databasePort"`
	Name     string `json:"databaseName"`
	SslMode  string `json:"databaseSslMode"`
	MaxConns int32  `json:"maxConns"`
}

type FilesConfig struct {
	TempFolder string `json:"tempFolder"`
}

type EstuaryConfig struct {
	UploadEndpoint string `json:"uploadEndpoint"`
	Token          string `json:"token"`
}

func CheckValidConfig(config Config, meta toml.MetaData) error {

	mainKeys := reflect.ValueOf(config)
	typeOfMainKeys := mainKeys.Type()

	for i := 0; i < mainKeys.NumField(); i++ {

		currentCatName := strings.ToUpper(typeOfMainKeys.Field(i).Name)

		// Check the category is defined
		if !meta.IsDefined(currentCatName) {
			return fmt.Errorf(currentCatName + " missing the config file")
		}

		// Cast the subKeys to all existing config structs
		// If cast works, check the keys exists in config.toml file
		subKeys := reflect.ValueOf(mainKeys.Field(i).Interface())

		for j := 0; j < subKeys.NumField(); j++ {

			if subKeys.Field(j).Type().Kind() == reflect.String {
				if subKeys.Field(j).Len() == 0 {
					return fmt.Errorf("Invalid value on index " + strconv.Itoa(j) + " of category " + currentCatName + " in config file")
				}
			}

			if subKeys.Field(j).Type().Kind() == reflect.Uint16 {
				if subKeys.Field(j).IsZero() {
					return fmt.Errorf("Invalid value on index " + strconv.Itoa(j) + " of category " + currentCatName + " in config file")
				}
			}
		}
	}

	return nil
}
