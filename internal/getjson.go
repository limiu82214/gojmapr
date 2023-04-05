package getjson

import (
	"encoding/json"
)

func Unmarshal(jsonString []byte, e interface{}) error {
	err := json.Unmarshal(jsonString, e)
	if err != nil {
		return err
	}
	return nil
}
