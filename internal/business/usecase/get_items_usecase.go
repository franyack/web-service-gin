package usecase

import (
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
)

type (
	GetItemsInterface interface {
		GetItems() (items []*domain.Item, err error)
	}

	getItemsInterface struct{}
)

func NewGetItemsUserCase() GetItemsInterface {
	return &getItemsInterface{}
}

func (getItems *getItemsInterface) GetItems() (items []*domain.Item, err error) {
	itemsRepository := gateway.NewItemsRepository()
	items, err = itemsRepository.GetItems()
	if err != nil {
		return nil, err
	}
	return items, err
}
