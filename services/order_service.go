package services

import "CS362-G7-FoodDelivery/models"

type OrderService interface {
	CreateOrder(customerID int, restaurantID int, deliveryAddr models.Geo) (models.Order, error)
	GetOrderDetail(orderID int) (models.Order, error)
}
