package models

type Payment struct {
	ID            uint8   `json:"id"`
	PaymentMethod string  `json:"paymentMethod"`
	Price         float64 `json:"price"`
}
