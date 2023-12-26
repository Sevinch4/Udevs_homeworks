package main

import (
	"fmt"
	"math/rand"
)

//task-1
func task1(n int,){
	slice := []int{}
	for i := 0; i < n; i++ {
		slice = append(slice, i+1)
	}
	fmt.Println(slice)
	max :=0
	min:=0
	//var a,b int
	for i := 0; i < len(slice); i++ {
		if max < slice[i]{
			max = i
		}
		if min > slice[i]{
			min = i
		}
	}
	slice[min],slice[max]=slice[max],slice[min]
	fmt.Println(slice)
}

//task-2
func tubS(){
	//var slice []int 
	slice := []int{1,11,7,3,5,6,12,9,10,13}
	 tub := []int{}
	for i := 0; i < 10; i++ {
		fmt.Printf("Input %d element: ",i+1)
		var num int
		fmt.Scan(&num)
		slice = append(slice, num)
	}
	fmt.Println(slice)
	
	for i := 0; i < len(slice); i++ {
		if slice[i] == 1{
			continue
		}
		count := 0
		for j := 1; j <= slice[i]/2; j++ {
			if slice[i]%j==0{
				count++
			}
		}
		
		if count <= 1{
			tub = append(tub, slice[i])
		}
		
	}
	fmt.Println(tub)
	fmt.Println("soni ",len(tub))
	
}
//task-3
func massiv(n int){
	array := make([]int , n) 
	for i := 0; i < n; i++ {
		fmt.Printf("Input %d element: ",i+1)
		fmt.Scan(&array[i])
	}
	fmt.Println(array)
	sumIndex :=0
	sumElements :=0
	for j := 0; j < len(array); j++ {
		sumIndex +=j
		sumElements +=array[j]
	}
	if sumIndex < sumElements{
		fmt.Println("INDEX")
	}else if sumIndex > sumElements{
		fmt.Println("VALUES")
	}else{
		fmt.Println("WOW")
	}
}
//task-4
func randomS(n int){
	//rand.Seed(time.Now().UnixNano())
	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(26+12)-12)
	}
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		if slice[i]<0{
			slice[i]=0
		}else{
			slice[i]=1
		}
	}
	fmt.Println(slice)
}

//task-5
func juft(n int){
	var slice []int 
	//var result string = "+"
	for i := 0; i < n; i++ {
		slice= append(slice, rand.Intn(35-14)+14)
	}
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		if slice[i]%2==0 {
			 fmt.Print(" + ")
		}else{
			fmt.Print(slice[i]," ")
		}
	}
	fmt.Println()
	
}
//task-6
func toq(n int){
	var s []int 
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(5+46)-46)
	}
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0{
			fmt.Print(s[i]," ")
		}else{
			fmt.Print(" _ ")
		}
	}
	fmt.Println()
}
//task-7
func swap1(n int){
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(25-5)-5)
	}
	fmt.Println(s)
	s[0],s[len(s)-1]=s[len(s)-1],s[0]
	fmt.Println(s)
	//s[len(s)-1]=s[0]
}

//task-8
func swap2(n int){
	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(25-5)-5)
	}
	fmt.Println(slice)
	slice[0],slice[len(slice)-1]=slice[len(slice)-1],slice[0]
	fmt.Println(slice)
}
//task-9
func swap3(n int){
	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(8+15)-15)
	}
	fmt.Println(slice)
	element := slice[len(slice)-1]
	copy(slice[1:],slice[:len(slice)-1])
	slice[0]=element
	fmt.Println(slice)
}

//task-10
func swapK(n,k int){
	var s []int 
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(21-8)+8)
	}
	fmt.Println(s)
	k %= len(s)
	temp := append(s[:0:0],s[:k]...)
	copy(s,s[k:])
	copy(s[len(s)-k:],temp)
	fmt.Println(s)
}
//task-11
func swapL(n,l int){
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(70+11)-11)
	}
	fmt.Println(s)
	l %= len(s)
	temp := append(s[:0:0],s[l:]...)
	copy(s,s[:l])
	copy(s[len(s)-l:],temp)
	fmt.Println(s)
}
func main(){
	//task-1
	// var n int
	// fmt.Print("Input num: ")
	// fmt.Scan(&n)
	// task1(n)

	//task-2
	// var t int
	// fmt.Print("Input num: ")
	// fmt.Scan(&t)
	tubS()

	//task-3
	// var n1 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n1)
	// massiv(n1)

	//task-4
	// var n2 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n2)
	// randomS(n2)

	//task-5
	// var n3 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n3)
	// juft(n3)

	//task-6
	// var a int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&a)
	// toq(a)

	//task-7
	// var n4 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n4)
	// swap1(n4)

	//task-8
	// var n5 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n5)
	// swap2(n5)

	//task-9
	// var n6 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n6)
	// swap3(n6)

	//task-10
	// var n7 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n7)
	// var k int
	// fmt.Print("Input k: ")
	// fmt.Scan(&k)
	// swapK(n7,k)

	//task-11
	// var n8 int 
	// fmt.Print("Input num: ")
	// fmt.Scan(&n8)
	// var l int
	// fmt.Print("Input k: ")
	// fmt.Scan(&l)
	// swapK(n8,l)







}