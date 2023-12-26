package dealar

import (
	"fmt"
	"math/rand"
	"main/store/product"
)
type Dealer struct{}

func (d Dealer) ProvideProduct(productName string,budget,quantity int)(product.Product,bool){
	originalprice := generateProductPrice(1,20)*1_000
	var temp int = 0

	if budget < originalprice * quantity{
		temp = budget/originalprice
		if temp < 1{
			fmt.Println("CASE",originalprice)
			return product.Product{},false
		}
	}else{
		temp = quantity
	}

	return product.Product{
		Name: productName,
		Quantity: temp,
		Price: originalprice*2,
		Originalprice: originalprice,
	},true
}
func generateProductPrice(min,max int)int{
	return rand.Intn(max-min) + min
}