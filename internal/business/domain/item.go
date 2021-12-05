package domain

import (
	"fmt"
	"time"
)

type Item struct {
	ID          int        `json:"id,omitempty"`
	Title       string     `json:"title"`
	Price       float32    `json:"price"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}

func (item *Item) CheckItem() error {

	if item.Title == "" {
		return fmt.Errorf("the field 'title' is empty or doesn't exists")
	}

	if item.Price <= 0 {
		return fmt.Errorf("the field 'price' cannot be less or equal to 0 or doesn't exists")
	}

	return nil
}
