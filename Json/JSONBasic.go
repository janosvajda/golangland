package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Simple string array
type simpleArray struct {
	Array []string
}

func main() {

	jsondat := &simpleArray{Array: []string{"one", "two", "three", "four", "five"}}
	encjson, err := json.Marshal(jsondat)

	if err == nil {
		fmt.Println(string(encjson))
	} else {
		log.Println(err)
	}

}
