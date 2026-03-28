package repositories

import "CS362-G7-FoodDelivery/models"

type CartRepository interface {
	SaveCart(cart models.Cart) error
	FindCartByCustomerID(customerID int) (models.Cart, error)
	UpdateCartItem(orderItemID int, quantity int, specialInstructions string) error
	DeleteCartItem(orderItemID int) error
}
