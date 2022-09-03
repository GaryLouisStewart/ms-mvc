package domain

import (
	"fmt"
	"net/http"

	"github.com/GaryLouisStewart/ms-mvc/utils"
)

var (
	elements = map[int64]*Element{
		1: {Id: 1, Name: "Hydrogen", AtomicMass: 1.00784, MeltingPoint: -259.2, BoilingPoint: -252.9, DiscoveryDate: 1766},
	}
)

func GetElement(elementId int64) (*Element, *utils.MsMvcError) {
	if element := elements[elementId]; element != nil {
		return element, nil
	}
	return nil, &utils.MsMvcError{
		Message:    fmt.Sprintf("element %v was not found", elementId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
