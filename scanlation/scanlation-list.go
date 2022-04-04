package scanlation

import (
	"encoding/json"
	"fmt"
	http "mangadex/http"
	utils "mangadex/utils"
)

type scanlationListRequest struct {
	Limit           int                 `json:"limit"`
	Offset          int                 `json:"offset"`
	Ids             []string            `json:"ids"`
	Name            string              `json:"name"`
	FocusedLanguage string              `json:"focusedLanguage"`
	Includes        []string            `json:"includes"`
	Order           scanlationListOrder `json:"order"`
}

type scanlationListOrder struct {
	Name          string `json:"name"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	FollowedCount string `json:"followedCount"`
	Relevance     string `json:"relevance"`
}

type scanlationListResponse struct {
	Result string               `json:"result"`
	Data   []scanlationListData `json:"data"`
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Total  int                  `json:"total"`
}

type scanlationListData struct {
	Id         string                   `json:"id"`
	Type       string                   `json:"type"`
	Attributes scanlationListAttributes `json:"attributes"`
}

type scanlationListAttributes struct {
	Name            string   `json:"name"`
	Website         string   `json:"website"`
	Discord         string   `json:"discord"`
	ContactEmail    string   `json:"contactEmail"`
	Description     string   `json:"description"`
	Twitter         string   `json:"twitter"`
	FocusedLanguage []string `json:"focusedLanguage"`
	Locked          bool     `json:"locked"`
	Official        bool     `json:"official"`
	Inactive        bool     `json:"inactive"`
	PublishDelay    string   `json:"publishDelay"`
	Version         int      `json:"version"`
	CreatedAt       string   `json:"createdAt"`
	UpdatedAt       string   `json:"updatedAt"`
}

func ScanlationList(
	limit int, offset int, ids []string, name string,
	focusedLang string, includes []string) {

	url := utils.RootUrl() + "/group"

	// create POST body
	var bodyOrder scanlationListOrder
	bodyOrder.CreatedAt = "asc"

	var body scanlationListRequest
	body.Limit = limit
	body.Offset = offset
	body.Ids = ids
	body.Name = name
	body.FocusedLanguage = focusedLang
	body.Includes = includes
	body.Order = bodyOrder

	json_data, err := json.Marshal(&body)
	utils.HandleError(err)

	response := http.Request("GET", url, json_data, "")

	var responseObject scanlationListResponse
	json.Unmarshal(response, &responseObject)

	var it int
	for it < len(responseObject.Data) {
		fmt.Printf("{SCANLATION-LIST} [%d]: %s\n", it, responseObject.Data[it].Attributes.Name)
		it++
	}
}
