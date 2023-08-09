package models

type Product struct {
	Product_ID   int
	Product_Name string
	Price        int64
}

type Cart struct {
	Cart_ID int
	Price   int
}
