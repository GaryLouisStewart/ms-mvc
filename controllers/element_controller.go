package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GaryLouisStewart/ms-mvc/services"
)

func GetElement(resp http.ResponseWriter, req *http.Request) {
	elementId, err := strconv.ParseInt(req.URL.Query().Get("element_id"), 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("element_id must be a number"))
		return
	}

	element, err := services.GetElement(elementId)
	if err != nil {
		// handle the error and return to the client
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	// return the user to the client

	jsonValue, _ := json.Marshal(element)
	resp.Write(jsonValue)

}
