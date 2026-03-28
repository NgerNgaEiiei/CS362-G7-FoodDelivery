package repositories

import "CS362-G7-FoodDelivery/models"

type OrderRepository interface {
	SaveOrder(order models.Order) error
	FindOrderByID(orderID int) (models.Order, error)
}
