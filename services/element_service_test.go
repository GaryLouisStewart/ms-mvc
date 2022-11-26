package services

import (
	"net/http"
	"testing"

	"github.com/GaryLouisStewart/ms-mvc/domain"
	"github.com/GaryLouisStewart/ms-mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	elementDaoMock elementsDaoMock

	getElementFunction func(elementId int64) (*domain.Element, *utils.MsMvcError)
)

func init() {
	domain.ElementDao = &elementsDaoMock{}
}

type elementsDaoMock struct{}

func (m *elementsDaoMock) GetElement(elementId int64) (*domain.Element, *utils.MsMvcError) {
	return getElementFunction(elementId)
}

func TestGetElementNotFoundInDatabase(t *testing.T) {
	getElementFunction = func(elementId int64) (*domain.Element, *utils.MsMvcError) {
		return nil, &utils.MsMvcError{
			StatusCode: http.StatusNotFound,
			Message:    "element 0 does not exist",
		}
	}
	element, err := ElementService.GetElement(1)
	assert.Nil(t, element)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "element 0 does not exist", err.Message)
}

func TestGetElementNoError(t *testing.T) {

	getElementFunction = func(elementId int64) (*domain.Element, *utils.MsMvcError) {
		return &domain.Element{
			Id: 1,
		}, nil
	}
	element, err := ElementService.GetElement(1)
	assert.Nil(t, err)
	assert.NotNil(t, element)
	assert.EqualValues(t, 1, element.Id)
}
