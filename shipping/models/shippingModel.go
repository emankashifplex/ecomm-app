package models

type Order struct {
	ID             int
	RecipientAddr  string
	Weight         float64
	ShippingOption string
}

func CalculateShippingCost(order Order) float64 {
	// Calculate shipping cost based on order details
	shippingCost := order.Weight * 0.5 // Example calculation

	switch order.ShippingOption {
	case "standard":
		shippingCost *= 1.0
	case "expedited":
		shippingCost *= 1.5
	case "overnight":
		shippingCost *= 2.0
	}

	return shippingCost
}
