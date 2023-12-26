package category_repo

import (
	"database/sql"
	"fmt"
	"main/market/categories"
	"os"
	"text/tabwriter"
)

type CategoryRepo struct {
	db *sql.DB
}

func NewCategory(db *sql.DB) CategoryRepo {
	return CategoryRepo{db: db}
}

// insert
func (c CategoryRepo) InsertCategory(category categories.Category) error {
	if _, err := c.db.Exec(`insert into category (id,name,created_at,updated_at) values ($1,$2,$3,$4)`, &category.ID, &category.Name, &category.Created_at, &category.Updated_at); err != nil {
		return err
	}
	return nil
}

// update
func (c CategoryRepo) UpdateCategoryByID(category categories.Category) error {
	if _, err := c.db.Exec(`update category set name = $1,updated_at = $2 where id = $3`, &category.Name, &category.Updated_at, &category.ID); err != nil {
		return err
	}
	return nil
}

// get list
func (c CategoryRepo) GetList() error {
	rows, err := c.db.Query(`select * from category`)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t Name\t Created_at\t Updated_at\n")
	for rows.Next() {
		cat := categories.Category{}
		if err = rows.Scan(&cat.ID, &cat.Name, &cat.Created_at, &cat.Updated_at); err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", cat.ID, cat.Name, cat.Created_at, cat.Updated_at)
	}
	w.Flush()
	return err
}

// get by id
func (c CategoryRepo) GetCategoryByID(id string) error {
	cat := categories.Category{}
	row := c.db.QueryRow(`select * from category where id = $1`, id)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t Name\t Created_at\t Updated_at\n")

	if err := row.Scan(&cat.ID, &cat.Name, &cat.Created_at, &cat.Updated_at); err != nil {
		return err
	}

	fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", cat.ID, cat.Name, cat.Created_at, cat.Updated_at)
	w.Flush()
	return nil
}

// delete
func (c CategoryRepo) DeleteCategory(id string) error {
	if _, err := c.db.Exec(`delete from category where id = $1`, id); err != nil {
		return err
	}
	return nil
}
