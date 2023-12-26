package tickets_repository

import (
	"database/sql"
	"fmt"
	"main/autopolet/tickets"
	"os"
	"text/tabwriter"
)

type TicketsRepository struct {
	db *sql.DB
}

func NewTicket(db *sql.DB) TicketsRepository {
	return TicketsRepository{
		db: db,
	}
}

// add ticket
func (t TicketsRepository) AddTicket(ticket tickets.Ticket) error {

	if _, err := t.db.Exec(`insert into ticket values($1,$2,$3,$4)`, &ticket.ID, &ticket.Departure, &ticket.Destination, &ticket.Date); err != nil {
		return err
	}
	return nil
}

// get by id
func (t TicketsRepository) GetByID(id string) error {
	ticket := tickets.Ticket{}
	row := t.db.QueryRow(`select * from ticket where id = $1`, id)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t From\t To\t Date\n")
	if err := row.Scan(&ticket.ID, &ticket.Departure, &ticket.Destination, &ticket.Date); err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", ticket.ID, ticket.Departure, ticket.Destination, ticket.Date)
	w.Flush()
	return nil
}

// get list
func (t TicketsRepository) GetList() error {
	rows, err := t.db.Query(`select * from ticket`)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t From\t To\t Date\n")

	for rows.Next() {
		t := tickets.Ticket{}
		if err = rows.Scan(&t.ID, &t.Departure, &t.Destination, &t.Date); err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", t.ID, t.Departure, t.Destination, t.Date)
		//ticket = append(ticket, t)
	}

	w.Flush()

	return nil
}

// update by id
func (t TicketsRepository) UpdateByID(ticket tickets.Ticket) error {
	if _, err := t.db.Exec(`update ticket set departure = $1,destination = $2,date = $3 where id = $4`, &ticket.Departure, &ticket.Destination, &ticket.Date, &ticket.ID); err != nil {
		return err
	}
	return nil
}

// delete ticket
func (t TicketsRepository) DeleteTicket(id string) error {
	if _, err := t.db.Exec(`delete from ticket where id = $1`, id); err != nil {
		return err
	}
	return nil
}
