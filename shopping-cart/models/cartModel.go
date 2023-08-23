package models

type CartItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

type Cart struct {
	UserID string
	Items  map[string]CartItem // map[ProductID]CartItem
}
