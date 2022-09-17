package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GaryLouisStewart/ms-mvc/utils"
)

var (
	elements = map[int64]*Element{
		1: {Id: 1, Name: "Hydrogen", AtomicMass: 1.00784, MeltingPoint: -259.2, BoilingPoint: -252.9, DiscoveryDate: 1766},
	}

	ElementDao elementDaoInterface
)

func init() {
	ElementDao = &elementDao{}
}

type elementDaoInterface interface {
	GetElement(int64) (*Element, *utils.MsMvcError)
}

type elementDao struct{}

func (u *elementDao) GetElement(elementId int64) (*Element, *utils.MsMvcError) {
	log.Println("we are accessing the database")
	if element := elements[elementId]; element != nil {
		return element, nil
	}
	return nil, &utils.MsMvcError{
		Message:    fmt.Sprintf("element %v does not exist", elementId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
