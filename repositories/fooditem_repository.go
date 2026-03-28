package repositories

import "CS362-G7-FoodDelivery/models"

type FoodItemRepository interface {
	FindFoodItemByID(foodItemID int) (models.FoodItem, error)
}
