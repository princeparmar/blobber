package models

import (
	"encoding/json"

	"gorm.io/datatypes"
)

// ToJSON convert object to JSON
func ToJSON(obj interface{}) datatypes.JSON {

	buf, err := json.Marshal(obj)

	if err != nil {
		return nil
	}

	return datatypes.JSON(string(buf))

}

// FromJSON parse object from JSON
func FromJSON(buf *datatypes.JSON, dest interface{}) error {

	retBytes, err := buf.Value()
	if err != nil {
		return err
	}
	if retBytes != nil {
		return json.Unmarshal(retBytes.([]byte), dest)
	}
	return nil
}
