package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpGet(url string) *http.Response {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	return resp
}

func httpPost(url string, data []byte) *http.Response {

	// json_data, err := json.Marshal(data)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	return resp
}

func processResponse(response *http.Response) []byte {

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
