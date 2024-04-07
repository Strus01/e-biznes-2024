package models

type Product struct {
	ID    uint8   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
