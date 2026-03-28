package models

type Order struct {
	OrderID      int
	CustomerID   int
	RestaurantID int
	Status       string
	Items        []OrderItem
	TotalPrice   int
}
