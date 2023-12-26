package store

import (
	"fmt"
	_ "main/store/basket"
	"main/store/costumer"
	"main/store/dealar"
	"main/store/product"
	"main/store/repo"
	"os"
	"text/tabwriter"
)

type Store struct {
	Dealer     dealar.Dealer
	Repository repo.Repository
	Budget     int
	Profit     int
}

func NewStore(repository repo.Repository) Store {
	return Store{
		Repository: repository,
		Budget:     100_000,
		Profit:     0,
	}
}

func (s *Store) StartSell(user costumer.Costumer) {
	//productSellRequest := products.GetProductInfo()
	productSellRequest := product.GetProductInfo()
	p, exists := s.Repository.Search(productSellRequest.Name)
	if !exists {
		fmt.Printf("We do not have this products\nWe will bring %s in the next time\n", productSellRequest.Name)
		p, success := s.Dealer.ProvideProduct(productSellRequest.Name, s.Budget, productSellRequest.Quantity)
		if !success {
			fmt.Println("We will buy !!!!!!")
			return
		}

		s.Repository.Products.AddProduct(p)
		s.Budget -= p.Originalprice * p.Quantity
		return
	}
	if user.Cash < p.Price*productSellRequest.Quantity {
		fmt.Println("You don't have enough cash")
		return
	}
	if p.Quantity < productSellRequest.Quantity {
		fmt.Printf("We do not have enough %s products, left %d\n", productSellRequest.Name, p.Quantity)
		return
	}

	s.Repository.TakeProduct(productSellRequest.Name, productSellRequest.Quantity)

	user.AddProduct(product.Product{
		Name:     p.Name,
		Quantity: productSellRequest.Quantity,
		Price:    p.Price,
	})
	//fmt.Println("person basket: ",person.Basket.UserBasket())
	user.Basket.UserBasket()
	s.Profit += productSellRequest.Quantity * (p.Price - p.Originalprice)
}

func (s *Store) PrintStats() {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', 0)
	fmt.Fprintln(w, "Name\tQuantity\tPrice\tOriginalPrice\t")
	for _, product := range s.Repository.Products {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\n", product.Name, product.Quantity, product.Price, product.Originalprice)
	}
	fmt.Fprintf(w, "\t\t\t\tBudget:%d\n", s.Budget)
	fmt.Fprintf(w, "\t\t\t\tProfit:%d\n", s.Profit)
	w.Flush()

}
