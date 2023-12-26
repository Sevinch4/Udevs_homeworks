package costumer

import (
	"fmt"
	"main/store/basket"
	"main/store/product"
)
type Costumer struct{
	Name string
	Cash int 
	Basket basket.Basket
}
func (c *Costumer) AddProduct(p product.Product){
	c.Basket.ProductList.AddProduct(p)
	for _,p := range c.Basket.ProductList{
		c.Basket.Total += p.Price * p.Quantity
	}
}
func NewCostumer(name string,cash int,Basket basket.Basket )Costumer{
	return Costumer{
		Name: name,
		Cash: cash,
		Basket: Basket,
	}
}
func GetCostumerInfo()(name string,cash int){
	fmt.Print("enter your name: ")
	fmt.Scan(&name)
	fmt.Println()
	fmt.Print("enter your cash: ")
	fmt.Scan(&cash)
	fmt.Println()
	return
}