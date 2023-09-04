# Users-service:

## Register
curl -X POST -H "Content-Type: application/json" -d '{"username": "yourusername", "password": "yourpassword"}' http://localhost:8080/register
## Login
curl -X POST -H "Content-Type: application/json" -d '{"username": "yourusername", "password": "yourpassword"}' http://localhost:8080/login
## User Profile 
curl http://localhost:8080/profile/yourusername
## Get Cart Endpoint
curl http://localhost:8080/get-cart
## Get Product Endpoint
curl http://localhost:8080/get-product-info/yourproductID

# Product-catalog

## Create a new product
curl -X POST -d '{"name": "Sample Product", "description": "This is a sample product.", "price": 19.99, "availability": true}' -H "Content-Type: application/json" http://localhost:8081/products
## Get product by ID
curl http://localhost:8081/products/{productID}
## Search Products
curl http://localhost:8081/products/search?query=sample
## Filter Products
curl http://localhost:8081/products/filter?minPrice=10&maxPrice=50&availability=true

# Shopping-cart
## Add item to cart
curl -X POST -H "Content-Type: application/json" -H "User-Id: 123" -d '{"product_id": 1, "quantity": 2, "price": 9.99}' http://localhost:8082/add-to-cart
## Remove item from cart
curl -X POST -H "Content-Type: application/json" -H "User-Id: 123" -d '1' http://localhost:8082/remove-from-cart
## Update cart item
curl -X POST -H "Content-Type: application/json" -H "User-Id: 123" -d '{"product_id": 1, "quantity": 3}' http://localhost:8082/update-cart-item
## Get cart
curl "http://localhost:8082/get-cart?user_id=123"

# Order

## Create Order
curl -X POST -H "Content-Type: application/json" -d '{"Product":"Product Name", "Quantity": 2}' http://localhost:8083/orders
## Get order
curl http://localhost:8083/orders/1

# Shipping
## Calculate shipping cost
curl -X POST -H "Content-Type: application/json" -d '{
    "ID": 1,
    "RecipientAddr": "456 Elm St",
    "Weight": 5.0,
    "ShippingOption": "standard"
}' http://localhost:8084/calculate_shipping_cost

TEST_DATABASE_URL=""postgres://eman:123@localhost:5432/product_data"