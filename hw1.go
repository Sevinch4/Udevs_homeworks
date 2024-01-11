package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Println("homework1")

	//task
	slice := []int{}
	uniqueNums := make(map[int]bool)

	for i := 0; i <= 50; i++ {
		r := rand.Intn(50)
		//for _, j := range slice {
		if !uniqueNums[r] {
			uniqueNums[r] = true
			slice = append(slice, r)
		} else {
			i--
		}
		//}

	}
	fmt.Println(slice)

	//	result := make([]int )

}
