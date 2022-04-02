package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	utils "mangadex/utils"
	"net/http"
)

func processResponse(response *http.Response) []byte {
	responseData, err := ioutil.ReadAll(response.Body)
	utils.HandleError(err)
	return responseData
}

func Request(method string, url string, data []byte, token string) []byte {

	fmt.Printf("{HTTP}: %s at %s\n", method, url)

	// create request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	utils.HandleError(err)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Authorization", token)

	// send request
	client := &http.Client{}
	resp, err := client.Do(req)
	utils.HandleError(err)

	fmt.Printf("{HTTP}: %s\n", resp.Status)
	// fmt.Printf("{HTTP}: %s\n", resp.Header)
	defer resp.Body.Close()
	return processResponse(resp)
}
