package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostRequest(url string, data interface{}) int {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	// POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(bytes))
	return resp.StatusCode
}
