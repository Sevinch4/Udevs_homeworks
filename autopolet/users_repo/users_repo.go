package users_repo

import (
	"database/sql"
	"fmt"
	"main/autopolet/users"
	"os"
	"text/tabwriter"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUser(db *sql.DB) UsersRepo {
	return UsersRepo{
		db: db,
	}
}

// add user
func (u UsersRepo) AddUSer(us users.User) error {
	if _, err := u.db.Exec(`insert into users values($1,$2,$3,$4,$5,$6)`, &us.ID, &us.First_name, &us.Last_name, &us.Email, &us.Phone, &us.Ticket_id); err != nil {
		return err
	}
	return nil
}

// get by id
func (u UsersRepo) GetByID(id string) error {
	us := users.User{}
	row := u.db.QueryRow(`select * from users where id = $1`, id)

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t Full name\t Email\t Phone\t Ticket_id\n")
	if err := row.Scan(&us.ID, &us.First_name, &us.Last_name, &us.Email, &us.Phone, &us.Ticket_id); err != nil {
		return err
	}
	fulName := fmt.Sprintf("%s %s", us.First_name, us.Last_name)
	fmt.Fprintf(w, "%s\t %s\t %s\t %s\t %s\n", us.ID, fulName, us.Email, us.Phone, us.Ticket_id)
	w.Flush()
	return nil
}

// get list
func (u UsersRepo) GetList() error {
	rows, err := u.db.Query(`select * from users`)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "ID\t Full name\t Email\t Phone\t Ticket_id\n")
	for rows.Next() {
		u := users.User{}
		if err = rows.Scan(&u.ID, &u.First_name, &u.Last_name, &u.Email, &u.Phone, &u.Ticket_id); err != nil {
			return err
		}
		fulName := fmt.Sprintf("%s %s", u.First_name, u.Last_name)
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\t %s\n", u.ID, fulName, u.Email, u.Phone, u.Ticket_id)
	}
	w.Flush()
	return nil

}

// update by id
func (u UsersRepo) UpdateByID(user users.User) error {
	if _, err := u.db.Exec(`update users set first_name = $1,last_name = $2,email = $3,phone = $4,ticket_id = $5 where id = $6`,
		&user.First_name, &user.Last_name, &user.Email, &user.Phone, &user.Ticket_id, &user.ID); err != nil {
		return err
	}
	return nil
}

// delete user
func (u UsersRepo) DeleteUser(id string) error {
	if _, err := u.db.Exec(`delete from users where id = $1`, id); err != nil {
		return err
	}
	return nil
}
