package models

type Pizza struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Price  float64  `json:"price"`
	Review []Review `json:"reviews"`
}
