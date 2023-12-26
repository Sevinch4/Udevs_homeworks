package main

import (
	"fmt"
	"math/rand"
)

//const chances = 3

func main() {
	player := Player{}
	game := Game{}

	name, age, favouriteNumber, chances := getplayerName()
	newPlayer := player.NewPlayer(name, age, favouriteNumber, chances)
	newGame := game.NewGame(newPlayer)
	newGame.StartGame()
}

type Game struct {
	RandomNumber int
	Player       Player
	tubson       bool
}

func tubSon(n int) bool {
	var isTrue bool
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			isTrue = false
		} else {
			isTrue = true
		}
	}
	return isTrue
}
func (g Game) NewGame(player Player) Game {
	if g.Player.chances > 5 {
		g.RandomNumber = rand.Intn(1000)
	} else {
		g.RandomNumber = rand.Intn(100)
	}
	return Game{
		RandomNumber: g.RandomNumber,
		Player:       player,
	}
}
func (g Game) StartGame() {
	fmt.Printf("Welcome %s\n", g.Player.name)
	fmt.Println("This is guessing game")
	g.tubson = tubSon(g.RandomNumber)
	if g.tubson {
		fmt.Println("Random number is prime!")
	}
	for i := 0; i < g.Player.chances; i++ {
		var n int
		fmt.Printf("%d chance enter number: ", i+1)
		fmt.Scan(&n)
		if n == g.RandomNumber {
			if g.RandomNumber == n {
				if g.Player.favouriteNumber == g.RandomNumber {
					fmt.Println("You won Random number is your favourite num: ", g.Player.favouriteNumber)
				}
			fmt.Println("You won! you found random number in ", i+1, "tries,random num is: ", g.RandomNumber)
			return
		} else {
			fmt.Println("Incorrect")
			if n > g.RandomNumber {
				fmt.Println("Random number lower than your enter number.")
			} else {
				fmt.Println("Random number greater than your enter number.")
			}
		}
	}

	fmt.Println("You lost!\nThe random number was", g.RandomNumber)
}
}

type Player struct {
	name            string
	age             int
	favouriteNumber int
	chances         int
}

func (p Player) NewPlayer(name string, age int, favouriteNumber int, chances int) Player {
	return Player{
		name:            name,
		age:             age,
		favouriteNumber: favouriteNumber,
		chances:         chances,
	}
}
func getplayerName() (string, int, int, int) {
	var (
		name            string
		age             int
		favouriteNumber int
		chances         int
	)
	fmt.Print("enter your name: ")
	fmt.Scan(&name)
	fmt.Print("enter your age: ")
	fmt.Scan(&age)
	fmt.Print("enter your favourite number: ")
	fmt.Scan(&favouriteNumber)
	fmt.Print("enter how many chances you want: ")
	fmt.Scan(&chances)

	return name, age, favouriteNumber, chances
}
