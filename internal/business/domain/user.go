package domain

import "time"

type Item struct {
	ID          string    `json:"id"`
	SiteID      string    `json:"site_id"`
	Title       string    `json:"title"`
	Price       float32   `json:"price"`
	DateCreated time.Time `json:"date_created"`
}
