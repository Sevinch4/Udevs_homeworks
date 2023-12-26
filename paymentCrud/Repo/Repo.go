package Repo

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"main/paymentCrud/payment"
	"os"
	"text/tabwriter"
)

type Repository struct {
	db *sql.DB
}

func NewPayment(db *sql.DB) Repository {
	return Repository{db: db}

}

// select
func (r Repository) Select(payment_type string) error {
	rows, err := r.db.Query(`select * from payments where payment_type = $1`, payment_type)
	if err != nil {
		return nil
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t Payment type\t Date\t Amount\n")
	for rows.Next() {
		p := payment.Payment{}
		if err = rows.Scan(&p.Id, &p.Payment_type, &p.Date, &p.Amount); err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", p.Id, p.Payment_type, p.Date, p.Date)
	}
	w.Flush()
	return nil
}

// add
func (r Repository) AddPayment(p payment.Payment) error {
	p.Id = uuid.New()
	if _, err := r.db.Exec(`insert into payments (id,payment_type,date,amount) values ($1,$2,$3,$4)`, p.Id, p.Payment_type, p.Date, p.Amount); err != nil {
		return err
	}
	return nil
}

// update
func (r Repository) UpdateByID(p payment.Payment) error {
	if _, err := r.db.Exec(`update payments set payment_type = $1,date = $2 ,amount = $3 where id = $4`, p.Payment_type, p.Date, p.Amount, p.Id); err != nil {
		return err
	}
	return nil
}

// delete ticket
func (r Repository) DeleteTicket(id string) error {
	if _, err := r.db.Exec(`delete from payments where id = $1`, id); err != nil {
		return err
	}
	return nil
}
