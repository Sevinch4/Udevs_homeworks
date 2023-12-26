package main

import (
	"fmt"
)

func add(n int) {
	slice := make([]int, 0, 10)
	var a int
	f := 0
	for i := 1; i <= n; i++ {
		a = i
		if a%2 == 0 {
			a *= -1
		}
		slice = append(slice, a)
	}
	fmt.Println(slice)
	var x, y int

	for i := 0; i < len(slice); i++ {

		if slice[i] < 0 {
			f++

			if f == 2 {
				x = i
			}
			if f == 4 {
				y = i
				break
			}
		}
	}
	sum := 0

	for i := x + 1; i < y; i++ {
		sum += slice[i]
	}
	fmt.Println(sum)

}

// task-2
func array(arr []int, n int) {
	var (
		nol, musbat, manfiy, juft, toq int = 0, 0, 0, 0, 0
	)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			nol++
		} else if arr[i] > 0 {
			musbat++
		} else if arr[i] < 0 {
			manfiy++
		}
		if arr[i]%2 == 0 {
			juft++
		} else {
			toq++
		}
	}
	fmt.Printf("Musbatlar %dta,Manfiy %dta,Nollar %dta,Juft %d,Toq %d\n", musbat, manfiy, nol, juft, toq)
}

// task-3
func sliceAj(array []int) {
	juft := []int{}
	toq := []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			juft = append(juft, array[i])
		} else {
			toq = append(toq, array[i])
		}
	}
	fmt.Println(juft)
	fmt.Println(toq)
}

// task-4
func slice(slice1 []int) {
	juft := 1
	toq := 0
	for i := 0; i < len(slice1); i++ {
		if i%2 == 0 {
			juft *= slice1[i]
		} else {
			toq += slice1[i]
		}
	}
	fmt.Println("Juft (products): ", juft)
	fmt.Println("Toq(sum): ", toq)
}

// taks-5
func middle(slice2 []int) {
	min1 := slice2[0]
	min2 := slice2[1]
	i := 0
	for i = 0; i < len(slice2); i++ {
		if slice2[i] < min1 {
			min1 = slice2[i]
		}

		if slice2[i] < min2 && slice2[i] != min1 {
			min2 = slice2[i]
		}
	}

	fmt.Println("Second min: ", min2)
}

// task-6
func maxSlice(slice3 []int) {
	max := slice3[0]
	max_2 := slice3[1]
	i := 0
	for i = 0; i < len(slice3); i++ {
		if slice3[i] > max {
			max = slice3[i]
		}

		if slice3[i] > max_2 && slice3[i] != max {
			max_2 = slice3[i]
		}
	}
	fmt.Println("Second max: ", max_2)
}

// task-7
func swap(n int, slice []int) {
	for i := 0; i < n; i++ {
		slice = append(slice, i+1)
	}
	slice[0], slice[n-1] = slice[n-1], slice[0]
	// slice[0],slice[len(slice)-1]=slice[len(slice)-1],slice[0]
	fmt.Println(slice)
}

// task-8
func maxN(slice []int) {
	max := 0
	f := 0
	var a int
	//var num int
	for i := 0; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
			a = i
		}
	}
	for j := a + 1; j < len(slice); j++ {
		f++
	}

	fmt.Printf("Maax sondan keyn %dta son bor\n", f)

}

// task-10
func minMax(sliceM []int) {
	max := sliceM[0]
	min := sliceM[1]
	var a, b int
	for i := 0; i < len(sliceM); i++ {
		if max < sliceM[i] {
			max = sliceM[i]
			a = i
		}
		if min > sliceM[i] {
			min = sliceM[i]
			b = i
		}
	}
	count := 0
	if a < b {
		for i := a + 1; i < b; i++ {
			count++
		}
	} else {

	}
}

//func main() {
//task-1
// var n int
// fmt.Print("Input num: ")
// fmt.Scan(&n)
// add(n)

//task-2
// var n1 = 10
// var arr []int = []int{0, -1, 2, -3, 4, 0, 5, -6, 7, 0}
// array(arr, n1)

//task-3
// var arr1 []int = []int{2, 3, 5, 4, 12, 78, 90, 45, 76, 99, 87}
// sliceAj(arr1)

//task-4
// slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// slice(slice1)

//task-5
// slice2 := []int{10,3, 4,0, 6, 7, 8, 5, 90}
// middle(slice2)

//task-6
// var slice3 = []int{3, 4, 2, 7, 8, 90, 6, 10, 14}
// maxSlice(slice3)

//task-7
// var n2 int
// fmt.Print("Input num:")
// fmt.Scan(&n2)
// slice4 := []int{}
// swap(n2,slice4)

//task-8
// var n3 int
// fmt.Print("Input num:")
//fmt.Scan(&n3)
// slice4 := []int{1, 4, 56, 3, 4, 78, 12, 3, 4, 5, 6}
// maxN(slice4)

//task-9
//task-10

//}
