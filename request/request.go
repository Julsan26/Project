package request

type Car struct {
	Make     string `json:"Make"`
	Model    string `json:"Model"`
	Package  string `json:"Package"`
	Color    string `json:"Color"`
	Year     string `json:"Year"`
	Catagory string `json:"Catagory"`
	Mileage  string `json:"Mileage"`
	Price    string `json:"Price"`
	ID       string `json:"ID"`
}
