package persist

import (
	"encoding/json"
)

func SaveModel(model interface{}, path string) error {
	payload, err := json.Marshal(model)
	if err != nil {
		return err
	}
	return saveToFile(path, payload)
}
