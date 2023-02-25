package main

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
)

func main() {

	jsonArrayFile, err := os.Open("../plaintext/plaintext_array.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonArrayFile.Close()

	listByteValue, _ := ioutil.ReadAll(jsonArrayFile)

    var list []interface{}
    json.Unmarshal([]byte(listByteValue), &list)

    fmt.Println(list)

	jsonObjectFile, err := os.Open("../plaintext/plaintext_object.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonObjectFile.Close()

	objectByteValue, _ := ioutil.ReadAll(jsonObjectFile)

    var obj map[string]interface{}
    json.Unmarshal([]byte(objectByteValue), &obj)

    fmt.Println(obj)

}

