package product_repo

import (
	"database/sql"
	"fmt"
	"main/market/products"
	"os"
	"text/tabwriter"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProduct(db *sql.DB) ProductRepo {
	return ProductRepo{db: db}
}

// insert
func (p ProductRepo) InsertProduct(prod products.Product) error {
	if _, err := p.db.Exec(`insert into product (id,name,price,category_id,created_at,updated_at) values($1,$2,$3,$4,$5,$6)`,
		&prod.ID, &prod.Name, &prod.Price, &prod.Category_id, &prod.Created_at, &prod.Updated_at); err != nil {
		return err
	}
	return nil
}

// update
func (p ProductRepo) UpdateProductByID(prod products.Product) error {
	if _, err := p.db.Exec(`update product set name = $1,price = $2,category_id = $3,updated_at = $4 where id = $5`,
		&prod.Name, &prod.Price, &prod.Category_id, &prod.Updated_at, &prod.ID); err != nil {
		return err
	}
	return nil
}

// get list
func (p ProductRepo) GetList() error {
	rows, err := p.db.Query(`select * from product`)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "Id\t Name\t Price\t Category_id\t Created_at\t Updated_at\n")

	for rows.Next() {
		pr := products.Product{}
		if err = rows.Scan(&pr.ID, &pr.Name, &pr.Price, &pr.Category_id, &pr.Created_at, &pr.Updated_at); err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\t %s\t %d\t %s\t %s\t %s\n", pr.ID, pr.Name, pr.Price, pr.Category_id, pr.Created_at, pr.Updated_at)
	}

	w.Flush()

	return nil
}

// get by id
func (p ProductRepo) GetProductByID(id string) error {
	pr := products.Product{}
	row := p.db.QueryRow(`select * from product where id = $1`, id)

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "Id\t Name\t Price\t Category_id\t Created_at\t Updated_at\n")

	if err := row.Scan(&pr.ID, &pr.Name, &pr.Price, &pr.Category_id, &pr.Created_at, &pr.Updated_at); err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\t %s\t %d\t %s\t %s\t %s\n", pr.ID, pr.Name, pr.Price, pr.Category_id, pr.Created_at, pr.Updated_at)
	w.Flush()

	return nil
}

// delete
func (p ProductRepo) DeleteProduct(id string) error {
	if _, err := p.db.Exec(`delete from product where id = $1`, id); err != nil {
		return err
	}
	return nil
}
