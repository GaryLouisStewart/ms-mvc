package services

import (
	"github.com/GaryLouisStewart/ms-mvc/domain"
)

func GetElement(elementId int64) (*domain.Element, error) {
	return domain.GetElement(elementId)
}
