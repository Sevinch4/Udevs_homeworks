package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/google/uuid"
	"main/person/user"
)

type Repo struct {
	db *sql.DB
}

func NewUser(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

// add data
func (r Repo) AddUser(u user.User) error {
	//user := user.User{
	//	Id:         id,
	//	First_name: first_name,
	//	Last_name:  last_name,
	//	Email:      email,
	//}
	if _, err := r.db.Exec(`insert into users values($1,$2,$3,$4)`, &u.Id, &u.First_name, &u.Last_name, &u.Email); err != nil {
		return err
	}
	return nil
}

// get by id
func (r Repo) GetByID(id int) (user.User, error) {
	u := user.User{}
	row := r.db.QueryRow(`select * from users where id = $1`, id)
	if err := row.Scan(&u.Id, &u.First_name, &u.Last_name, &u.Email); err != nil {
		fmt.Println("error is while selecting by id", err)
	}
	return u, nil
}

// get list
func (r Repo) GetList() ([]user.User, error) {
	users := []user.User{}
	rows, err := r.db.Query(`select * from users`)
	if err != nil {
		fmt.Println("error is while selecting data", err)
	}
	for rows.Next() {
		user := user.User{}
		if err = rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// delete by id
func (r Repo) DeleteByID(id int) error {
	if _, err := r.db.Exec(`delete from users where id = $1`, id); err != nil {
		return err
	}
	return nil
}
