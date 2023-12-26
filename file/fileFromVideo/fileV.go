package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewDecoder and Decode
//
//	type Car struct {
//		Brand string `json: "brand"`
//		Year  int    `json: "year"`
//		Model string `json: "model"`
//		Power int    `json: "power"`
//	}
type Product struct {
	Name          string `json: "name"`
	Quantity      int    `json: "quantity"`
	Price         int    `json: "price"`
	OriginalPrice int    `json: "original_price"`
}
type Basket struct {
	Products []Product
	Total    int `json: "total"` //omitempty agar json file da bomasa pustoy valuesini ciqaradi
}
type User struct {
	Name   string `json: "name"`
	Cash   int    `json: "cash"`
	Basket Basket
}
type Animal struct {
	Name  string `json: "name"`
	Type  string `json: "type"`
	Sound string `json: "sound"`
}

func main() {
	//animal example
	a := []Animal{}
	file, err := os.Open("JsonFile.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(file).Decode(&a); err != nil {
		panic(err)
	}
	fmt.Println(a)
	for _, animal := range a {
		if animal.Type == "wild" {
			fmt.Println(animal)
		}
	}

	//store sample
	// file,err := os.Open("JsonFile.json")
	// if err != nil{
	// 	panic(err)
	// }
	// person := User{}
	// if err := json.NewDecoder(file).Decode(&person); err != nil{
	// 	panic(err)
	// }
	// fmt.Println(person)

	//create file
	// file,err := os.Create("File1.txt")//RDWR - ochb ketadi qayta run time qganda
	// if err != nil{
	// 	fmt.Println("error while creating file",err.Error())
	// 	return
	// }

	//open file
	// file, err := os.Open("File.txt") //only open this file and checks existence of file is in this directory or not(is eror)
	// if err != nil {
	// 	fmt.Println("error while opening file", err.Error())
	// 	return
	// }
	// defer file.Close()

	//file read I
	// b := make([]byte,4)
	// n,err := file.Read(b)
	// if err != nil{
	// 	fmt.Println("error while reading file by file read",err.Error())
	// 	return
	// }
	// fmt.Println(n,string(b))

	//file read II
	// b,err := os.ReadFile("File.txt")
	// if err != nil{
	// 	fmt.Println("error while reading file by os.readfile",err.Error())
	// 	return
	// }
	// fmt.Println(string(b))

	//file read III
	// b,err := io.ReadAll(file)
	// if err != nil{
	// 	fmt.Println("error while io.readAll file",err.Error())
	// 	return
	// }
	// fmt.Println(string(b))

	//file read line by line
	// 	products := []Product{}
	// 	scanner := bufio.NewScanner(file)
	// 	for scanner.Scan(){//returns bool, data oqib bogunca,kursor bn oqidi
	// 		p := Product{}
	// 		line := scanner.Text()//string qataradi
	// 		slice := strings.Split(line," ")//slice ni " " probel boyica element qilib saqlab beradi
	// 		quantity,_ := strconv.Atoi(slice[1])//stringni int ga convert qiladi
	// 		price,_ := strconv.Atoi(slice[2])
	// 		originalP,_ := strconv.Atoi(slice[3])
	// 		p.Name = slice[0]
	// 		p.Quantity = quantity
	// 		p.Price = price
	// 		p.OriginalPrice = originalP
	// 		products = append(products, p)//bitta slice ga yigadi
	// 		//fmt.Println(slice)
	// 		//same thing split and fields
	// 		// slice1 := strings.Fields(line)
	// 		// fmt.Println(slice1)
	// 	}
	// 	fmt.Println(products)//bitda slicega yigilgan datani chiqaradi

	//NEWDECDER AND DECODE
	// c := Car{}
	// file,err := os.Open("JsonFile.json")//open json file
	// if err != nil{
	// 	fmt.Println("error while opening file",err.Error())
	// }
	// defer file.Close()//program tugaganda fileni close qivomiz
	// if err := json.NewDecoder(file).Decode(&c); err != nil{//newDecoder file
	// 	panic(err)
	// }
	// fmt.Println(c)

}
