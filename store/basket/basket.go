package basket

import (
	"fmt"
	"main/store/product"
	"os"
	"text/tabwriter"
)

type Basket struct{
	ProductList product.ProductList
	Total int

}
func NewBasket()Basket{
	return Basket{}
}
func (b Basket) UserBasket(){
	w := tabwriter.NewWriter(os.Stdout,1,8,1,'\t',0)
	fmt.Fprintln(w,"Name\tQuantity\tPrice\t")
	for _,basket := range b.ProductList{
		fmt.Fprintf(w,"%s\t%d\t%d\t\n",basket.Name,basket.Quantity,basket.Price)
	}
	fmt.Fprintf(w,"\t\t\t\tTotal: %d\n",b.Total)
	w.Flush()
}