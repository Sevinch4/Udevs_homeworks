package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"main/autopolet/report"
	"main/autopolet/tickets"
	"main/autopolet/tickets_repository"
	"main/autopolet/users"
	"main/autopolet/users_repo"
	"time"
)

func main() {
	db, err := ConnectinDB()
	if err != nil {
		fmt.Println("error is while connecting to database", err)
	}
	defer db.Close()

	ticRepo := tickets_repository.NewTicket(db)
	userRepo := users_repo.NewUser(db)
	repoReport := report.NewReport(db)

	var option int
	fmt.Print(`
                   1 - Ticket
                   2 - User
                   3 - Report
Choose option: 
`)
	fmt.Scan(&option)
	switch option {
	case 1: //ticket
		var t int
		fmt.Print(`
                      1 - add ticket
                      2 - get by id
                      3 - get list 
                      4 - update ticket by id
                      5 - delete ticket
choose option:
`)
		fmt.Scan(&t)
		switch t {
		case 1: //add ticket
			var (
				from string
				to   string
				date string
			)
			fmt.Print("input from: ")
			fmt.Scan(&from)
			fmt.Print("input to: ")
			fmt.Scan(&to)
			fmt.Print("input date (dd-mm-yyy): ")
			fmt.Scan(&date)
			Date, err := time.Parse("2006-01-02", date)
			if err != nil {
				fmt.Println("error is while parsing time", err)
			}
			ticket := tickets.Ticket{
				ID:          uuid.New(),
				Departure:   from,
				Destination: to,
				Date:        Date,
			}
			if err := ticRepo.AddTicket(ticket); err != nil {
				fmt.Println("error is while calling add ticket method", err)
				return
			}
			fmt.Println("ticket added successfully")
		case 2: //get by id
			var ticketID string
			fmt.Print("input ticket id: ")
			fmt.Scan(&ticketID)

			//get by id method
			err = ticRepo.GetByID(ticketID)
			if err != nil {
				fmt.Println("error is while calling get by id method", err)
				return
			}

		case 3: //get list
			err = ticRepo.GetList()
			if err != nil {
				fmt.Println("error is while calling get list method", err)
			}
		case 4: //update by id
			var (
				id    string
				fromU string
				toU   string
				dateU string
			)
			fmt.Print("input id: ")
			fmt.Scan(&id)
			fmt.Print("input from: ")
			fmt.Scan(&fromU)
			fmt.Print("input to: ")
			fmt.Scan(&toU)
			fmt.Print("input date (yyy-mm-dd): ")
			fmt.Scan(&dateU)

			DateU, err := time.Parse("2006-01-02", dateU)
			if err != nil {
				fmt.Println("error is while parsing time", err)
			}
			idP, _ := uuid.Parse(id)
			t := tickets.Ticket{
				ID:          idP,
				Departure:   fromU,
				Destination: toU,
				Date:        DateU,
			}
			if err = ticRepo.UpdateByID(t); err != nil {
				fmt.Println("error is while calling update by id method", err)
				return
			}
			fmt.Println("ticket updated successfully")
		case 5: //delete ticket
			var D_id string
			fmt.Print("input id: ")
			fmt.Scan(&D_id)

			//delete method
			if err = ticRepo.DeleteTicket(D_id); err != nil {
				fmt.Println("error is while callng method delete ticket")
			}
			fmt.Println("ticket deleted")
		}
	case 2: //user
		var u int
		fmt.Print(`
                      1 - add user
                      2 - get by id
                      3 - get list 
                      4 - update user by id
                      5 - delete user
choose option:
`)
		fmt.Scan(&u)
		switch u {
		case 1: //add user
			var (
				first_name string
				last_name  string
				email      string
				phone      string
				ticket_id  string
			)

			fmt.Print("input first name: ")
			fmt.Scan(&first_name)
			fmt.Print("input last name: ")
			fmt.Scan(&last_name)
			fmt.Print("input email: ")
			fmt.Scan(&email)
			fmt.Print("input phone: ")
			fmt.Scan(&phone)
			fmt.Print("input ticket_id: ")
			fmt.Scan(&ticket_id)

			user := users.User{
				ID:         uuid.New(),
				First_name: first_name,
				Last_name:  last_name,
				Email:      email,
				Phone:      phone,
				Ticket_id:  ticket_id,
			}
			//add method
			if err = userRepo.AddUSer(user); err != nil {
				fmt.Println("error is while calling add user method", err)
				return
			}
			//fmt.Println(user)
			fmt.Println("user added successfully")
		case 2: //get by id
			var idU string
			fmt.Print("input id: ")
			fmt.Scan(&idU)

			//get by id method
			err = userRepo.GetByID(idU)
			if err != nil {
				fmt.Println("error is while calling user repo method", err)
				return
			}
		case 3: //get list
			err = userRepo.GetList()
			if err != nil {
				fmt.Println("error is while calling get list method", err)
				return
			}
		case 4: //update by id
			var U_ID string
			var (
				first_name string
				last_name  string
				email      string
				phone      string
				ticket_id  string
			)
			//ticketA, _ := uuid.Parse(ticket_id)
			fmt.Print("input id: ")
			fmt.Scan(&U_ID)
			fmt.Print("input first name: ")
			fmt.Scan(&first_name)
			fmt.Print("input last name: ")
			fmt.Scan(&last_name)
			fmt.Print("input email: ")
			fmt.Scan(&email)
			fmt.Print("input phone: ")
			fmt.Scan(&phone)
			fmt.Print("input ticket_id: ")
			fmt.Scan(&ticket_id)
			id, _ := uuid.Parse(U_ID)
			user := users.User{
				ID:         id,
				First_name: first_name,
				Last_name:  last_name,
				Email:      email,
				Phone:      phone,
				Ticket_id:  ticket_id,
			}
			if err = userRepo.UpdateByID(user); err != nil {
				fmt.Println("error is while calling update method", err)
				return
			}
			fmt.Println("user updated")
		case 5: //delete user
			var D_id string
			fmt.Print("input id: ")
			fmt.Scan(&D_id)

			if err = userRepo.DeleteUser(D_id); err != nil {
				fmt.Println("error is while calling delete method", err)
				return
			}
			fmt.Println("user deleted")
		}
	case 3: //report
		var (
			R_from string
			R_to   string
		)
		fmt.Print("from: ")
		fmt.Scan(&R_from)
		fmt.Print("to: ")
		fmt.Scan(&R_to)

		//report
		if err = repoReport.Report(R_from, R_to); err != nil {
			fmt.Println("error is while calling report method", err)
		}
	}
}

func ConnectinDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		"host = localhost port = 5432 password = 1218 database = autopolet sslmode = disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
