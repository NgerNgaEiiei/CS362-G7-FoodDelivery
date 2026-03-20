package models

type Geo struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Address string  `json:"address,omitempty"`
}
