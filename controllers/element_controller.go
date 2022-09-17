package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/GaryLouisStewart/ms-mvc/services"
	"github.com/GaryLouisStewart/ms-mvc/utils"
)

func GetElement(c *gin.Context) {
	elementId, err := strconv.ParseInt(c.Param("element_id"), 10, 64)
	if err != nil {
		apiErr := &utils.MsMvcError{
			Message:    "element_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondWithError(c, apiErr)
		return
	}

	element, apiErr := services.ElementService.GetElement(elementId)
	if apiErr != nil {
		// handle the error and return to the client
		utils.RespondWithError(c, apiErr)
		return
	}

	// return the user to the client
	utils.RespondContentType(c, http.StatusOK, element)

}
