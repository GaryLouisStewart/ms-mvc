package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetElementNoElementFound(t *testing.T) {
	element, err := GetElement(0)

	assert.Nil(t, element, "not expecting an element with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "element 0 does not exist", err.Message)
}

func TestElementNoError(t *testing.T) {
	element, err := GetElement(1)
	assert.Nil(t, err)
	assert.NotNil(t, element)
	assert.EqualValues(t, 1, element.Id)
	assert.EqualValues(t, "Hydrogen", element.Name)
	assert.EqualValues(t, 1.00784, element.AtomicMass)
	assert.EqualValues(t, -259.2, element.MeltingPoint)
	assert.EqualValues(t, -252.9, element.BoilingPoint)
	assert.EqualValues(t, 1766, element.DiscoveryDate)

}
