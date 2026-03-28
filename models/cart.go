package models

type Cart struct {
	CartID     int
	CustomerID int
	Items      []OrderItem
	TotalPrice int
}
