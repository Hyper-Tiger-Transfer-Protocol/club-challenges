//FIXME - Need to make it so file ia actually parsed
package main

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
)

func main() {

	jsonFile, err := os.Open("../plaintext/plaintext.json")

	if err != nil {
		fmt.Println(err)
	}

	jsonData, err := json.Marshal(jsonFile)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}
fmt.Printf("json data: %s\n", jsonData)


	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    fmt.Println(result[])
}

