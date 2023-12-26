package product

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	List = ProductList{}
)

func (l *ProductList) ProdList() {
	//list := []ProductList{}
	file, err := os.Open("List.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(file).Decode(&List); err != nil {
		panic(err)
	}
}

//	type List struct{
//		Name string `json: "name"`
//		Price int `json: "price"`
//		Originalprice int `json: "original_price"`
//		Quantity int `json: "quantity"`
//	}
type Product struct {
	Name          string
	Price         int
	Originalprice int
	Quantity      int
}
type ProductSellRequest struct {
	Name     string
	Quantity int
	//Basket []Basket
}

func NewProduct(name string, price, quantity int) Product {
	return Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

type ProductList []Product

func (p *ProductList) AddProduct(product Product) {
	*p = append(*p, product)
}
func (p *ProductList) RemoveProduct(index int) {
	*p = append((*p)[:index], (*p)[index+1:]...)
}
func GetProductInfo() ProductSellRequest {
	var (
		productName string
		quantity    int = 0
	)

	fmt.Print("Enter products name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter products quantity: ")
	fmt.Scan(&quantity)

	fmt.Println()

	return ProductSellRequest{
		Name:     productName,
		Quantity: quantity,
	}
}
