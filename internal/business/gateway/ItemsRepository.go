package gateway

import "example/web-service-gin/internal/business/domain"

var itemsRepository ItemsRepository

func RegisterItemsRepository(repository ItemsRepository) {
	itemsRepository = repository
}

func NewItemsRepository() ItemsRepository {
	if itemsRepository == nil {
		panic("itemsRepository not found")
	}
	return itemsRepository
}

type ItemsRepository interface {
	GetItems() ([]*domain.Item, error)
	GetItemById(itemID string) (*domain.Item, error)
	AddItem(item *domain.Item) error
}
