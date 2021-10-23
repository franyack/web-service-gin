package repository

import (
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"fmt"
	"time"
)

func NewMySqlItemsRepository() gateway.ItemsRepository {
	return &mysqlItemsRepository{}
}

type mysqlItemsRepository struct {
}

func (repository *mysqlItemsRepository) AddItem(item *domain.Item) error {
	return nil
}

func (repository *mysqlItemsRepository) GetItemById(itemID string) (*domain.Item, error) {
	if itemID != "1234" {
		return nil, apierrors.NewNotFoundApiError(fmt.Sprintf("item_id %s was not found", itemID))
	}
	now := time.Now()
	return &domain.Item{
		ID:          "1234",
		SiteID:      "ARG",
		Title:       "Harry Potter and the Philosopher’s Stone",
		Price:       102.34,
		DateCreated: &now,
	}, nil
}

func (repository *mysqlItemsRepository) GetItems() ([]*domain.Item, error) {
	now := time.Now()
	itemOne := domain.Item{
		ID:          "1234",
		SiteID:      "ARG",
		Title:       "Harry Potter and the Philosopher’s Stone",
		Price:       102.34,
		DateCreated: &now,
	}
	itemTwo := domain.Item{
		ID:          "12345",
		SiteID:      "ARG",
		Title:       "Harry Potter and The Chamber of Secrets",
		Price:       102.34,
		DateCreated: &now,
	}
	var items []*domain.Item
	items = append(items, &itemOne)
	items = append(items, &itemTwo)
	return items, nil
}
