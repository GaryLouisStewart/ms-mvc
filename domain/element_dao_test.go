package domain

import (
	"net/http"
	"testing"

	"github.com/GaryLouisStewart/ms-mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	elementDaoMock elementsDaoMock

	getElementFunction func(elementId int64) (*Element, *utils.MsMvcError)
)

func init() {
	ElementDao = &elementsDaoMock{}
}

type elementsDaoMock struct{}

func (m *elementsDaoMock) GetElement(elementId int64) (*Element, *utils.MsMvcError) {
	return getElementFunction(elementId)
}

func TestGetElementNoElementFound(t *testing.T) {

	getElementFunction = func(elementId int64) (*Element, *utils.MsMvcError) {
		return nil, &utils.MsMvcError{
			StatusCode: http.StatusNotFound,
			Message:    "element 0 does not exist",
		}
	}

	element, err := DomainService.GetElement(0)
	assert.Nil(t, element, "not expecting an element with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "element 0 does not exist", err.Message)
}

func TestElementNoError(t *testing.T) {
	getElementFunction = func(elementId int64) (*Element, *utils.MsMvcError) {
		return &Element{
			Id:            1,
			Name:          "Hydrogen",
			AtomicMass:    1.00784,
			MeltingPoint:  -259.2,
			BoilingPoint:  -252.9,
			DiscoveryDate: 1766,
		}, nil
	}
	element, err := DomainService.GetElement(1)
	assert.Nil(t, err)
	assert.NotNil(t, element)
	assert.EqualValues(t, 1, element.Id)
	assert.EqualValues(t, "Hydrogen", element.Name)
	assert.EqualValues(t, 1.00784, element.AtomicMass)
	assert.EqualValues(t, -259.2, element.MeltingPoint)
	assert.EqualValues(t, -252.9, element.BoilingPoint)
	assert.EqualValues(t, 1766, element.DiscoveryDate)

}
