package services

import "CS362-G7-FoodDelivery/models"

type CartService interface {
	AddItemToCart(customerID int, foodItemID int, quantity int, specialInstructions string) (models.Cart, error)
	GetCartByCustomer(customerID int) (models.Cart, error)
	UpdateCartItem(orderItemID int, quantity int, specialInstructions string) error
	RemoveCartItem(orderItemID int) error
}
