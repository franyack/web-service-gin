package domain

import (
	"fmt"
	"time"
)

type Item struct {
	ID          string     `json:"id"`
	SiteID      string     `json:"site_id"`
	Title       string     `json:"title"`
	Price       float32    `json:"price"`
	DateCreated *time.Time `json:"date_created,omitempty"`
}

func (item *Item) CheckItem() error {
	if item.ID == "" {
		return fmt.Errorf("the field 'id' is empty or doesn't exists")
	}
	if item.SiteID == "" {
		return fmt.Errorf("the field 'site_id' is empty or doesn't exists")
	}

	if item.Title == "" {
		return fmt.Errorf("the field 'title' is empty or doesn't exists")
	}

	if item.Price <= 0 {
		return fmt.Errorf("the field 'price' cannot be less or equal to 0 or doesn't exists")
	}

	return nil
}
