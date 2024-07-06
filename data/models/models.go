package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Size  string  `json:"size"`
	Color string  `json:"color"`
}

type Manufacturer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Bike struct {
	Product
	FrameNumber string   `json:"frameNumber"`
	Owner       Customer `json:"owner"`
}

type Customer struct {
	Id        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Address   Address `json:"address"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}
