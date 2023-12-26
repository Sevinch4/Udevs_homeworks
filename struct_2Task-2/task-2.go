package main

import (
	"fmt"
)

type Clients struct {
	ID     int
	Name   string
	basket []CardProducts
}
type CardProducts struct {
	ProductID int
	SizeID    int
	Quantity  int
}
type Product struct {
	ID   int
	Name string
	Size []Size
}
type Size struct {
	ID    int
	Name  string
	Price int
}

var clients = []Clients{
	{
		ID: 1, Name: "Akhmedov", basket: []CardProducts{
			{ProductID: 1, SizeID: 2, Quantity: 4},
			{ProductID: 4, SizeID: 1, Quantity: 3},
			{ProductID: 3, SizeID: 3, Quantity: 1},
		},
	},
	{
		ID: 1, Name: "Komilov", basket: []CardProducts{
			{ProductID: 2, SizeID: 1, Quantity: 1},
			{ProductID: 2, SizeID: 2, Quantity: 6},
			{ProductID: 5, SizeID: 1, Quantity: 5},
		},
	},
}
var product = []Product{
	{ID: 1, Name: "Pizza", Size: []Size{
		{ID: 1, Name: "25sm", Price: 30000},
		{ID: 2, Name: "30sm", Price: 45000},
		{ID: 3, Name: "35sm", Price: 70000},
	}},
	{ID: 2, Name: "Burger", Size: []Size{
		{ID: 1, Name: "Small", Price: 15000},
		{ID: 2, Name: "Medium", Price: 25000},
		{ID: 3, Name: "Big", Price: 30000},
	}},
	{ID: 3, Name: "Hot-Dog", Size: []Size{
		{ID: 1, Name: "Small", Price: 20000},
		{ID: 2, Name: "Medium", Price: 25000},
		{ID: 3, Name: "Big", Price: 30000},
	}},
	{ID: 4, Name: "Coca-Cola", Size: []Size{
		{ID: 1, Name: "0.5", Price: 5000},
		{ID: 2, Name: "1.0", Price: 8000},
		{ID: 3, Name: "1.5", Price: 10000},
	}},
	{ID: 5, Name: "Fanta", Size: []Size{
		{ID: 1, Name: "0.5", Price: 5000},
		{ID: 2, Name: "1.0", Price: 8000},
		{ID: 3, Name: "1.5", Price: 11000},
	}},
}

func main() {
	for _, valClient := range clients {
		sum := 0
		fmt.Printf("%s - ", valClient.Name)
		for _, valBasket := range valClient.basket {
			for _, valProduct := range product {
				if valBasket.ProductID == valProduct.ID {
					for _, val := range valProduct.Size {
						if valProduct.ID == val.ID {
							sum += val.Price * valBasket.Quantity
						}

					}
				}
			}
		}
		fmt.Printf(" total sum : %d\n", sum)

	}
	for _, valclients := range clients {
		productName := ""
		max := 0
		for _, valProduct := range product {
			for _, val := range valclients.basket {
				if valProduct.ID == val.ProductID {
					if max < val.Quantity {
						max = val.Quantity
						productName = valProduct.Name
					}
				}

			}
		}
		fmt.Println(productName, max)
	}
}