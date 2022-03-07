package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")

	if err == nil {
		fmt.Println(string(b))
	}

	return
}

func ConvertBetweenStruct(from interface{}, to interface{}) (err error, toStruct interface{}) {
	b, err := json.Marshal(from)

	if err != nil {
		return err, interface{}(nil)
	}

	err = json.Unmarshal(b, to)

	if err != nil {
		return err, interface{}(nil)
	}

	return nil, to
}
