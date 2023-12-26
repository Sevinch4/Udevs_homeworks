package main

import (
	"fmt"
	"main/person/repo"
	"main/person/user"
)

/*
	input 1: user ni hamma malumotini terminal da oqib olib db ga saqlah kk +++
	input 2: user ni id sini terminadan oqib olib osha userni hamma malumotini chiqarish kk +++
	input 3: userlarni chiqarish kk +++
	insert 4: userni id sini terminaldan oqib olib osha id orqali user ni delete qilish kk
*/
import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	//connection to database
	db, err := ConnectionDB()
	if err != nil {
		fmt.Println("error is while connecting", err)
	}
	defer db.Close()

	rep := repo.NewUser(db)

	var option int
	fmt.Print(`Choose option:
						1 - add user
						2 - get user by id
						3 - users list
						4 - delete user by id
`)
	fmt.Scan(&option)
	switch option {
	case 1:
		var (
			id         int
			first_name string
			last_name  string
			email      string
		)
		fmt.Print("Input id: ")
		fmt.Scan(&id)
		fmt.Print("Input first name: ")
		fmt.Scan(&first_name)
		fmt.Print("Input last name: ")
		fmt.Scan(&last_name)
		fmt.Print("Input email: ")
		fmt.Scan(&email)

		//add method
		us := user.User{
			Id:         id,
			First_name: first_name,
			Last_name:  last_name,
			Email:      email,
		}
		if err = rep.AddUser(us); err != nil {
			fmt.Println("error is while adding user", err)
		}
		fmt.Println("user successfully added")
	case 2:
		var id int
		fmt.Print("Input your ID: ")
		fmt.Scan(&id)

		//get by id method
		user, err := rep.GetByID(id)
		if err != nil {
			fmt.Println("error is while using getByID method", err)
		}
		fmt.Println(user)
	case 3:
		fmt.Println("Users list: ")
		//getList method
		users, err := rep.GetList()
		if err != nil {
			fmt.Println("error is while using getList method", err)
		}
		fmt.Println(users)
	case 4:
		var id int
		fmt.Print("Input your id: ")
		fmt.Scan(&id)

		//delete method
		if err = rep.DeleteByID(id); err != nil {
			fmt.Println("error is while calling deleteByID method in main", err)
		}
		fmt.Println("Deleted successfully")
	}

}

func ConnectionDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port =5432 password = 1218 database=users sslmode = disable")
	if err != nil {
		fmt.Println("error is while opening database", err)
	}
	return db, nil
}
