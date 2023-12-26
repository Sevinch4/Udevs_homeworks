package main

import "fmt"

// task
func logic(n int) {
	var s string = ""
	var k int
	count := 0
	for i := 0; i < n; i++ {
		n /= 10
		count++
	}
	if n > 0 && n < 10 {
		k = (n / 1) % 10
		switch k {
		case 1:
			s += "bir"
		case 2:
			s += "ikki"
		case 3:
			s += "uch"
		case 4:
			s += "to'rt"
		case 5:
			s += "besh"
		case 6:
			s += "olti"
		case 7:
			s += "yetti"
		case 8:
			s += "sakkiz"
		case 9:
			s += "to'qqiz"

			fmt.Println(s)
		}
	} else if n > 10 && n < 100 {
		k = (n / 10) % 10
		switch k {
		case 1:
			s += "o'n "
		case 2:
			s += "yigirma "
		case 3:
			s += "o'ttiz "
		case 4:
			s += "qirq "
		case 5:
			s += "ellik "
		case 6:
			s += "oltmish "
		case 7:
			s += "yetmish "
		case 8:
			s += "sakson "
		case 9:
			s += "to'qson "
		}
		k = (n / 1) % 10
		switch k {
		case 1:
			s += " bir "
		case 2:
			s += " ikki "
		case 3:
			s += " uch "
		case 4:
			s += " to'rt "
		case 5:
			s += " besh "
		case 6:
			s += " olti "
		case 7:
			s += " yetti "
		case 8:
			s += " sakkiz "
		case 9:
			s += " to'qqiz "

			fmt.Println(s)

		}
	} else if n > 99 && n < 1000 {
		k = (n / 100) % 10
		switch k {
		case 1:
			s += "bir yuz "
		case 2:
			s += "ikki yuz "
		case 3:
			s += "uch yuz "
		case 4:
			s += "to'rt yuz "
		case 5:
			s += "besh yuz "
		case 6:
			s += "olti yuz "
		case 7:
			s += "yetti yuz "
		case 8:
			s += "sakkiz yuz "
		case 9:
			s += "to'qqiz yuz "
		}
		k = (n / 10) % 10
		switch k {
		case 1:
			s += " o'n "
		case 2:
			s += " yigirma "
		case 3:
			s += " o'ttiz "
		case 4:
			s += " qirq "
		case 5:
			s += " ellik "
		case 6:
			s += " oltmish "
		case 7:
			s += " yetmish "
		case 8:
			s += " sakson "
		case 9:
			s += " to'qson "
		}
		k = (n / 1) % 10
		switch k {
		case 1:
			s += " bir "
		case 2:
			s += " ikki "
		case 3:
			s += " uch "
		case 4:
			s += " to'rt "
		case 5:
			s += " besh "
		case 6:
			s += " olti "
		case 7:
			s += " yetti "
		case 8:
			s += " sakkiz "
		case 9:
			s += " to'qqiz "
		}
		fmt.Println(s)

	}

}

func reshetka(n int) {

	result := ""

	for i := n; i > 0; i-- {
		for k := i; k < n; k++ {
			result += " "
		}

		for j := 1; j <= i; j++ {
			result += "#"
		}

		fmt.Println(result)

		result = ""

		/*result := ""

		for i := 0; i < n; i++ {
			for k := 0; k < n-i-1; k++ {
				result += " "
			}

			for j := 0; j <= i; j++ {
				result += "#"
			}

			fmt.Println(result)

			result = ""*/

	}
	//fmt.Println()
}

func main() {
	//HOMEWORK
	var number int
	fmt.Print("Input num: ")
	fmt.Scan(&number)
	logic(number)

	//PRACTICE
	// //task-1
	// var a, b, c int
	// fmt.Print("Input numbers: ")
	// fmt.Scan(&a, &b, &c)
	// sum(a, b, c)
	// //task-2
	// var x1, x2, y1, y2 float64
	// fmt.Print("Input coordinates: ")
	// fmt.Scan(&x1, &x2, &y1, &y2)
	// coordinates(x1, x2, y1, y2)
	// // //task-3
	// var r, r1, r2 float64
	// fmt.Print("Input a,b,c: ")
	// fmt.Scan(&r, &r1, &r2)
	// equation(r, r1, r2)
	// //task-4
	// var n int
	// fmt.Print("Input number: ")
	// fmt.Scan(&n)
	// fmt.Println(fib(n))
	// //task-5
	// var n1 int
	// fmt.Print("Input number : ")
	// fmt.Scan(&n1)
	// reshetka(n1)

	// //task-6
	// var n2 int
	// fmt.Print("Input number: ")
	// fmt.Scan(&n2)
	// digits(n2)

	// //task-7
	// var n3, n4 int
	// fmt.Print("Input two numbers: ")
	// fmt.Scan(&n3, &n4)
	// fmt.Println(add(n3, n4))

	// //taks-8
	// var f int
	// fmt.Print("Input nummber: ")
	// fmt.Scan(&f)
	// fmt.Println(fact(f))
	// //task-9
	// var N int
	// fmt.Print("Input number: ")
	// fmt.Scan(&N)
	// reverse(N)

}
