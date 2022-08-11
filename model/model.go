package model

type Car struct {
	Make     string `json:"Make"`
	Model    string `json:"Model"`
	Package  string `json:"Package"`
	Color    string `json:"Color"`
	Year     string `json:"Year"`
	Category string `json:"Category"`
	Mileage  string `json:"Mileage"`
	Price    string `json:"Price"`
	ID       int64  `json:"ID"`
}
