package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"main/paymentCrud/Repo"
	"main/paymentCrud/payment"
	"time"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println("error is while connecting to database", err)
	}
	defer db.Close()

	repo := Repo.NewPayment(db)

	var option int
	fmt.Print(`
                   1 - select
                   2 - add
                   3 - update
                   4 - delete
Choose option: 
`)
	fmt.Scan(&option)
	switch option {
	case 1:
		var payment_type string
		fmt.Print("input payment type: ")
		fmt.Scan(&payment_type)
		if err = repo.Select(payment_type); err != nil {
			fmt.Println("error is while selecting", err)
		}
	case 2:
		var (
			payment_type string
			amount       int
		)
		fmt.Print("input payment type: ")
		fmt.Scan(&payment_type)
		fmt.Print("input amount: ")
		fmt.Scan(&amount)
		p := payment.Payment{
			Payment_type: payment_type,
			Date:         time.Now(),
			Amount:       amount,
		}
		if err = repo.AddPayment(p); err != nil {
			fmt.Println("error is while add payment", err)
			return
		}
		fmt.Println("added")
	case 3: //update
		var (
			id           string
			payment_type string
			amount       int
		)
		fmt.Print("input id:")
		fmt.Scan(&id)
		fmt.Print("input payment type: ")
		fmt.Scan(&payment_type)
		fmt.Print("input amount: ")
		fmt.Scan(&amount)

		p_id, err := uuid.Parse(id)
		p := payment.Payment{
			Id:           p_id,
			Payment_type: payment_type,
			Date:         time.Now(),
			Amount:       amount,
		}
		if err != nil {
			fmt.Println("error is while parsing", err)
			return
		}
		if err = repo.UpdateByID(p); err != nil {
			fmt.Println("error is while update", err)
			return
		}
		fmt.Println("updated")
	case 4: //delete
		var id string
		fmt.Print("input id:")
		fmt.Scan(&id)
		if err = repo.DeleteTicket(id); err != nil {
			fmt.Println("error is while delete", err)
			return
		}
		fmt.Println("delete")
	}

}
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host = localhost port = 5432 password = 1218 database = payment sslmode = disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
