package models

type FoodItem struct {
	FoodItemID   int
	RestaurantID int
	FoodName     string
	Price        int
	Description  string
	IsAvailable  bool
}
