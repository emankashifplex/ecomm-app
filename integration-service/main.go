package main

func main() {
	// Simulate a user registration and login.
	userID := registerUser("eman", "123")
	sessionToken := loginUser("eman", "123")

	//Simulate product search
	productID := createProduct()
	searchProducts("bag", 8.0, 50.0, true)

	//Simulate adding to, removing from, and viewing cart
	addToCart(userID, productID, 2, 12.00)
	removeFromCart(userID, productID)

	//Simulate creating and getting an order
	orderID := createOrder(userID, sessionToken, 7, 2)
	getOrder(orderID)

	//Simulate calculating shipping cost
	calculateShippingCost(orderID)

	//Simulate getting confirmation email
	sendConfirmationEmail(userID, orderID)
}
