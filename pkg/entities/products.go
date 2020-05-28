package entities

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PriceInCents uint64 `json:"price"`
}
