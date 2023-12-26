package main

import "fmt"

type Client struct {
	ID     int
	Name   string
	Basket []BasketProducts
}
type BasketProducts struct {
	ProductID int
	Quantity  int
}
type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	client := []Client{
		{
			ID:   1,
			Name: "Qobil",
			Basket: []BasketProducts{
				{
					ProductID: 1,
					Quantity:  3,
				},
				{
					ProductID: 2,
					Quantity:  3,
				},
				{
					ProductID: 3,
					Quantity:  3,
				},
			},
		},
		{
			ID:   2,
			Name: "Steve",
			Basket: []BasketProducts{
				{
					ProductID: 3,
					Quantity:  6,
				},
			},
		},
	}
	product := []Product{
		{
			ID:    1,
			Name:  "bread",
			Price: 5000,
		},
		{
			ID:    2,
			Name:  "banana",
			Price: 20000,
		},
		{
			ID:    3,
			Name:  "yogurt",
			Price: 5000,
		},
		{
			ID:    4,
			Name:  "kfc",
			Price: 100000,
		},
	}
	for _, val := range client {
		sum := 0
		fmt.Print("Client name ", val.Name)
		for _, valBasket := range val.Basket {
			for _, p := range product {

				if valBasket.ProductID == p.ID {
					sum += p.Price * valBasket.Quantity
				}

			}
		}
		fmt.Print(" Total Sum ", sum, "\n")
	}
}
