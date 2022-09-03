package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GaryLouisStewart/ms-mvc/services"
	"github.com/GaryLouisStewart/ms-mvc/utils"
)

func GetElement(resp http.ResponseWriter, req *http.Request) {
	elementId, err := strconv.ParseInt(req.URL.Query().Get("element_id"), 10, 64)
	if err != nil {
		apiErr := &utils.MsMvcError{
			Message:    "element_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	element, apiErr := services.GetElement(elementId)
	if apiErr != nil {
		// handle the error and return to the client
		resp.WriteHeader(apiErr.StatusCode)

		resp.Write([]byte(apiErr.Message))
		return
	}

	// return the user to the client

	jsonValue, _ := json.Marshal(element)
	resp.Write(jsonValue)

}
