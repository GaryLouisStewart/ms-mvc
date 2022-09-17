package services

import (
	"github.com/GaryLouisStewart/ms-mvc/domain"
	"github.com/GaryLouisStewart/ms-mvc/utils"
	"net/http"
)

type itemsService struct{}

var (
	ItemService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.MsMvcError) {
	return nil, &utils.MsMvcError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
