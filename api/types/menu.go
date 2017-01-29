package types

type Menu struct {
	Pizzaria string `json:"pizzaria"`
	Pizzas   []Pizza `json:"pizzas"`
}

type Pizza struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int `json:"price"`
}
