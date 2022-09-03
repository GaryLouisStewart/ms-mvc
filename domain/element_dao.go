package domain

import (
	"fmt"
)

var (
	elements = map[int64]*Element{
		1: {Id: 1, Name: "Hydrogen", AtomicMass: 1.00784, MeltingPoint: -259.2, BoilingPoint: -252.9, DiscoveryDate: 1766},
	}
)

func GetElement(elementId int64) (*Element, error) {
	if element := elements[elementId]; element != nil {
		return element, nil
	}
	return nil, fmt.Errorf(fmt.Sprintf("element %v was not found", elementId))
}
