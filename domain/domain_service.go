package domain

import "github.com/GaryLouisStewart/ms-mvc/utils"

type elementService struct{}

var (
	DomainService elementService
)

func (u *elementService) GetElement(elementId int64) (*Element, *utils.MsMvcError) {
	return ElementDao.GetElement(elementId)
}
