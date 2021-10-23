package usecase

import (
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
)

type (
	AddItemInterface interface {
		AddItem(item *domain.Item) error
	}

	addItemInterface struct{}
)

func NewAddItemUserCase() AddItemInterface {
	return &addItemInterface{}
}

func (getItems *addItemInterface) AddItem(item *domain.Item) error {
	repository := gateway.NewItemsRepository()
	if err := repository.AddItem(item); err != nil {
		return err
	}
	return nil
}
