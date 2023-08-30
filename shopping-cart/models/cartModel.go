package models

// Structure for an item that can be added to the cart
type CartItem struct {
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

// Structure for the shopping cart
type Cart struct {
	UserID int        `json:"user_id"`
	Items  []CartItem `json:"items"`
	Total  float64    `json:"total"`
}

// Creeates a shopping cart instance with an empty list of items
func NewCart() *Cart {
	return &Cart{
		Items: []CartItem{},
	}
}

// AddItem adds a new item to the cart and updates the total cost
func (c *Cart) AddItem(productID int, quantity int, price float64) {
	item := CartItem{
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
	c.Items = append(c.Items, item)      // Add the item to the cart's items list.
	c.Total += price * float64(quantity) // Update the total cost.
}

// RemoveItem removes an item from the cart and updates the total cost
func (c *Cart) RemoveItem(productID int) {
	var updatedItems []CartItem
	for _, item := range c.Items {
		if item.ProductID != productID {
			updatedItems = append(updatedItems, item)
		} else {
			c.Total -= item.Price * float64(item.Quantity) // Subtract the item's cost from the total
		}
	}
	c.Items = updatedItems // Update the cart's items list
}

// UpdateQuantity updates the quantity of a specific item in the cart and adjusts the total cost
func (c *Cart) UpdateQuantity(productID int, newQuantity int) {
	for i, item := range c.Items {
		if item.ProductID == productID {
			c.Total -= item.Price * float64(item.Quantity) // Subtract the current cost
			c.Items[i].Quantity = newQuantity              // Update the quantity
			c.Total += item.Price * float64(newQuantity)   // Add the updated cost
			break
		}
	}
}
