package services

import (
	"github.com/GaryLouisStewart/ms-mvc/domain"
	"github.com/GaryLouisStewart/ms-mvc/utils"
)

func GetElement(elementId int64) (*domain.Element, *utils.MsMvcError) {
	return domain.GetElement(elementId)
}
