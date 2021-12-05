package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJson(path string, v interface{}) {
	// Open our jsonFile

	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, v)
}
