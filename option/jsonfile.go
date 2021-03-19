package option

import (
	"encoding/json"
	"log"
	"os"
)

func JsonLoad(filename string, v interface{}) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}
	log.Println(filename, "Found")
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Println(err)
	}
}

func JsonSave(filename string, v interface{}) {
	jsonTxt, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile(filename, jsonTxt, 0644)
	if err != nil {
		log.Println(err)
	}
}
