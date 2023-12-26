package repo
import (
	"main/store/product"
)
type Repository struct{
	Products product.ProductList

}

func NewRepository(products product.ProductList)Repository{
	return Repository{
		Products: products,
	}
}

func (r *Repository) Search(productName string) (product.Product,bool){
	for _,product := range r.Products{
		if product.Name == productName{
			return product,true
		}
	}
	return product.Product{},false
}

func (r *Repository) TakeProduct(productName string, quantity int){
	for index,product := range r.Products{
		if product.Name == productName{
			//basket := append(r,)
			r.Products[index].Quantity -= quantity
			if r.Products[index].Quantity == 0{
				r.Products.RemoveProduct(index)
			}
			return
		}
	}
}