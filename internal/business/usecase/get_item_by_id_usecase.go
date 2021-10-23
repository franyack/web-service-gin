package usecase

import (
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
)

type (
	GetItemByIdInterface interface {
		GetItemById(itemID string) (item *domain.Item, err error)
	}

	getItemByIdInterface struct{}
)

func NewGetItemByIdUserCase() GetItemByIdInterface {
	return &getItemByIdInterface{}
}

func (getItems *getItemByIdInterface) GetItemById(itemID string) (item *domain.Item, err error) {
	itemsRepository := gateway.NewItemsRepository()
	item, err = itemsRepository.GetItemById(itemID)
	if err != nil {
		return nil, err
	}
	return item, err
}
