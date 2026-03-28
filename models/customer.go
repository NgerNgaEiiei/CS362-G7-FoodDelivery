package models

type Customer struct {
	CustomerID      int
	CustomerName    string
	CustomerPhone   string
	CustomerAddress Geo
}
