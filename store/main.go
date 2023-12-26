package main

import (
	"fmt"
	"main/store/repo"
	"main/store/product"
	"main/store/store"
	"main/store/basket"
	"main/store/costumer"
)
const (
	StartShopCmd = iota + 1
	FinishShopCmd
)
const(
	BossCmd = iota + 1
	Costumer
	UserQuitCmd

)
//boss cmd
const(
	ReportCmd = iota + 1
	ListCmd
	BackCmd
	QuitCmd
)
const Password = "1234"
func main() {
	repository := repo.NewRepository(product.List)
	store1 := store.NewStore(repository)

	for true {
		var userCmd int 
		//var cmd int
		fmt.Print(`
		Enter command:
		1 - Boss
		2 - Costumer
		3 - Quit
`)
		fmt.Scan(&userCmd)

		switch userCmd {
		case BossCmd:
			var password string
			fmt.Print("enter password: ")
			fmt.Scan(&password)
			if password != Password{
				fmt.Println("Password is wrong")
				return
			}
			fmt.Println("Boss login")
			fmt.Print(`
		Enter command:
		1 - Report
		2 - Product List
		3 - Back
		4 - Quit
`)
			var bossCmd int
			fmt.Print("Enter command")
			fmt.Scan(&bossCmd)
			switch bossCmd{
			case ReportCmd:
				store1.PrintStats()
			case ListCmd:
				fmt.Println(product.List)
			case BackCmd:
			case QuitCmd:
			}
		case Costumer:
			b := basket.NewBasket()
			name, cash := costumer.GetCostumerInfo()
			user := costumer.NewCostumer(name,cash,b)

			var cmdUser int
	
			fmt.Println(`
		Enter command:
		1 - Start shopping
		2 - my basket
		3 - finish
		4 - back
		5 - quit
`)
			fmt.Scan(&cmdUser)
			switch cmdUser{
			case StartShopCmd:
				store1.StartSell(user)
			case FinishShopCmd:
				
				return
			}
		case UserQuitCmd:
			return
		default:
			fmt.Println("There is not such kind of command!!!")
		}
	}
}