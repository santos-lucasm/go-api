package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handleError(err error) {

	if err != nil {
		fmt.Printf("{ERROR}: %s", err.Error())
		log.Fatal(err)
	}
}

func processResponse(response *http.Response) []byte {

	responseData, err := ioutil.ReadAll(response.Body)
	handleError(err)
	return responseData
}

func httpRequest(method string, url string, data []byte, token string) []byte {

	fmt.Printf("{HTTP}: %s at %s\n", method, url)

	// create request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	handleError(err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	handleError(err)

	fmt.Printf("{HTTP}: %s\n", resp.Status)
	// fmt.Printf("{HTTP}: %s\n", resp.Header)
	defer resp.Body.Close()
	return processResponse(resp)
}
