package main

import "fmt"

func main(){
	var number int 
	fmt.Print("enter Number: ")
	fmt.Scan(&number)
	if number < 0{
		number = -1*number
		fmt.Print("-")
		numberToWord(number)
	}else{numberToWord(number)}
	
}

func numberToWord(n int){
	var(
		dig int = 0
		ten int = 1
		workNumber int = n
	)
	for workNumber != 0{
		workNumber = workNumber / 10
		dig++
	}
	for i := 1; i < dig; i++ {
		ten = ten * 10
	}
	for dig >= 1{
		var s string
		workNumber = n / ten 
		if dig % 3 == 2{
			switch workNumber{
			case 1:
				s = "o'n"
				fmt.Print(s)
			case 2:
				s = "yigirma"
				fmt.Print(s)
			case 3:
				s = "o'ttiz"
				fmt.Print(s)
			case 4:
				s = "qirq"
				fmt.Print(s)
			case 5:
				s = "ellik"
				fmt.Print(s)
			case 6:
				s = "oltmish"
				fmt.Print(s)
			case 7:
				s = "yetmish"
				fmt.Print(s)
			case 8:
				s = "sakson"
				fmt.Print(s)
			case 9:
				s= "to'qson"
				fmt.Print(s)
			}
			fmt.Print(" ")
		}else{
			switch workNumber{
			case 1:
				s = "bir"
				fmt.Print(s)
			case 2:
				s = "ikki"
				fmt.Print(s)
			case 3:
				s = "uch"
				fmt.Print(s)
			case 4:
				s = "to'rt"
				fmt.Print(s)
			case 5:
				s = "besh"
				fmt.Print(s)
			case 6:
				s = "olti"
				fmt.Print(s)
			case 7:
				s = "yetti"
				fmt.Print(s)
			case 8:
				s = "sakkiz"
				fmt.Print(s)
			case 9:
				s= "to'qqiz"
				fmt.Print(s)
			}
			fmt.Print(" ")
		}
		if dig % 3 == 0 && workNumber != 0{
			fmt.Print("yuz")
		}
		if workNumber != 0 {
			switch dig{
			case 4:
				s= "ming"
				fmt.Print(s)
			case 7:
				s = "million"
				fmt.Print(s)
				case 10:
					s = "trillion"
					fmt.Print(s) 
			}
			fmt.Print(" ")
		}
		n %= ten
		ten /= 10
		dig--
	}
	fmt.Println()
}