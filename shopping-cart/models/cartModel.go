package model

type CartItem struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Cart struct {
	UserID int        `json:"user_id"`
	Items  []CartItem `json:"items"`
	Total  float64    `json:"total"`
}

func NewCart() *Cart {
	return &Cart{
		Items: []CartItem{},
	}
}

func (c *Cart) AddItem(productID int, quantity int, price float64) {
	item := CartItem{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
	c.Items = append(c.Items, item)
	c.Total += price * float64(quantity)
}

func (c *Cart) RemoveItem(productID int) {
	var updatedItems []CartItem
	for _, item := range c.Items {
		if item.ProductID != productID {
			updatedItems = append(updatedItems, item)
		} else {
			c.Total -= item.Price * float64(item.Quantity)
		}
	}
	c.Items = updatedItems
}

func (c *Cart) UpdateQuantity(productID int, newQuantity int) {
	for i, item := range c.Items {
		if item.ProductID == productID {
			c.Total -= item.Price * float64(item.Quantity)
			c.Items[i].Quantity = newQuantity
			c.Total += item.Price * float64(newQuantity)
			break
		}
	}
}
