package option

import (
	"encoding/json"
	"os"

	"github.com/scutrobotlab/asuwave/internal/logger"
)

func JsonLoad(filename string, v interface{}) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}
	logger.Log.Println(filename, "Found")
	data, err := os.ReadFile(filename)
	if err != nil {
		logger.Log.Println(err)
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		logger.Log.Println(err)
	}
}

func JsonSave(filename string, v interface{}) {
	jsonTxt, err := json.Marshal(v)
	if err != nil {
		logger.Log.Println(err)
	}
	err = os.WriteFile(filename, jsonTxt, 0644)
	if err != nil {
		logger.Log.Println(err)
	}
}
