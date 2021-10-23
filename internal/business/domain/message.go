package domain

type Message struct {
	Message string `json:"message"`
	Item    *Item  `json:"item,omitempty"`
}
