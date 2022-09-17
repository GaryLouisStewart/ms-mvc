package services

import (
	"github.com/GaryLouisStewart/ms-mvc/domain"
	"github.com/GaryLouisStewart/ms-mvc/utils"
)

type elementService struct{}

var (
	ElementService elementService
)

func (u *elementService) GetElement(elementId int64) (*domain.Element, *utils.MsMvcError) {
	return domain.ElementDao.GetElement(elementId)
}
