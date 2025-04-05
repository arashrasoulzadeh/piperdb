package persist

import (
	"encoding/json"
)

func SaveModel(model interface{}, filename string) error {
	payload, err := json.Marshal(model)
	if err != nil {
		return err
	}
	return saveToFile(filename, payload)
}
