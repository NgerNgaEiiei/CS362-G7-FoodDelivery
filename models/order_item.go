package models

type OrderItem struct {
	OrderItemID         int
	FoodItemID          int
	Quantity            int
	UnitPrice           int
	SpecialInstructions string
}
