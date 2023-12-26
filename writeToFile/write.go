package main

import (
	"os"
)

// task
var m = map[string]string{
	"name":   "ism",
	"age":    "yosh",
	"hello":  "salom",
	"home":   "uy",
	"school": "maktab",
	"try":    "harakat",
	"horse":  "ot",
	"monkey": "maymun",
	"dog":    "kuchuk",
	"cat":    "mushuk",
}

func main() {
	file, err := os.OpenFile("File.txt", os.O_CREATE|os.O_RDWR, 06666) //openFile is open file or create file
	if err != nil {
		panic(err)
	}
	//file write I
	// n, err := file.WriteString("hello world")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n)

	//file write II
	// if _, err := file.Write([]byte("hello guys")); err != nil {
	// 	panic(err)
	// }

	//file write III
	// if _,err := io.WriteString(file,"salom dunyo"); err != nil{
	// 	panic(err)
	// }

	//file write IV
	// if err := os.WriteFile("File.txt",[]byte("salom"),0666); err !=nil{
	// 	panic(err)
	// }

	//task
	// file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// _, err = file.WriteString("english - uzb\n\n")
	// if err != nil{
	// 	panic(err)
	// }
	// c := 0
	// for i, v := range m {
	// 	c++
	// 	v += "\n"
	// 	str := fmt.Sprintf("%d. %s - %s",c,i,v)
	// 	if _, err := io.WriteString(file, str); err != nil {
	// 		panic(err)
	// 	}
	// }

}
