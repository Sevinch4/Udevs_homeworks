package report

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewReport(db *sql.DB) Repo {
	return Repo{db: db}
}

type Reports struct {
	From           string
	To             string
	First_name     string
	Last_name      string
	Customer_phone string
	Date           time.Time
}

func (r Repo) Report(from, to string) error {
	counter := 0
	rows, err := r.db.Query(`select t.departure as from, t.destination as to, first_name,last_name,phone as Customer_phone,t.date from users as u join ticket as t on u.ticket_id = t.id where t.departure = $1 and t.destination = $2`, &from, &to)
	if err != nil {
		return err
	}
	//tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "№\t From\t To\t Customer_Full_name\t customer_phone\t date\n")

	for rows.Next() {
		counter++
		r := Reports{}
		if err = rows.Scan(&r.From, &r.To, &r.First_name, &r.Last_name, &r.Customer_phone, &r.Date); err != nil {
			return err
		}
		formatedDate := r.Date.Format("January 02, 2006")
		fullName := fmt.Sprintf("%s %s", r.First_name, r.Last_name)
		fmt.Fprintf(w, "%d\t %s\t %s\t %s\t %s\t %s\n", counter, r.From, r.To, fullName, r.Customer_phone, formatedDate)
	}
	w.Flush()
	return nil
}
