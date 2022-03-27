package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpRequest(url string) *http.Response {

	resp, err := http.Get(url)
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
