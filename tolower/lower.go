package main

import "fmt"

func toLower(s string)string{
	result := ""
	for _,i := range s {
		if 'A' <= i && i <= 'Z' {
			result += string(i + 32)// ASCII "S" = 83, 83 + 32 = 115, ASCII 115 is "s"
		}else{
			result += string(i)
		}
	}
	return result
}
func main(){
	var s string
	fmt.Scan(&s)
	fmt.Println(toLower(s))
}