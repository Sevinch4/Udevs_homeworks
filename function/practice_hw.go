package main

import (
	"fmt"
	"math"
)

// PRACTICE:
// task-1
func sum(a, b, c int) {
	max := 0
	min := 0
	if a > max {
		max = a
	} else if b > max {
		max = b
	} else {
		max = c
	}
	if a < min {
		min = a
	} else if b < min {
		min = b
	} else {
		min = c
	}
	sum := max + min
	fmt.Println(sum)

}

// task-2
func coordinates(x1, x2, y1, y2 float64) {
	var z float64
	z = math.Sqrt(math.Pow((x2-x1), 2) + math.Pow((y2-y1), 2))
	fmt.Printf(" %.1f\n", z)
}

// task-3
func equation(a, b, c float64) {
	var d float64
	var root1 float64
	var root2 float64
	d = b*b - 4*a*c
	if d > 0 {
		root1 = (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
		root2 = (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
		fmt.Printf("root1 = %.1f, root2 = %.1f\n", root1, root2)
	} else if d == 0 {
		root1 = -b / 2 * a
		fmt.Printf("One root = %.1f\n", root1)
	} else {
		fmt.Println("No solution for this equation (d<0)!")
	}
}

// task-4
func fib(n int) int {
	f0 := 0
	f1 := 1
	if n == 0 {
		return f0
	} else if n == 1 {
		return f1
	}
	var fn int
	if n > 1 {
		fn = fib(n-1) + fib(n-2)

	}
	return fn
}

// task-5
// func reshetka(n int) {

// 	result := ""

// 	for i := n; i >0 ; i-- {
// 		for k := i; k < n; k++ {
// 			result += " "
// 		}

// 		for j := 0; j <= i; j++ {
// 			result += "#"
// 		}

// 		fmt.Println(result)

// 		result = ""

// 		/*result := ""

// 	for i := 0; i < n; i++ {
// 		for k := 0; k < n-i-1; k++ {
// 			result += " "
// 		}

// 		for j := 0; j <= i; j++ {
// 			result += "#"
// 		}

// 		fmt.Println(result)

// 		result = ""*/

// 	}
// 	//fmt.Println()
// }

// task-6
func digits(n int) {
	//digit := n
	digit := math.Log10(math.Abs(float64(n))) + 1
	fmt.Printf("%.d\n", int(digit))
}

// task-7
func add(n, n1 int) int {
	digit := 1
	for i := n1; i > 0; i /= 10 {
		digit *= 10
	}
	result := n*digit + n1
	return result
}

// task-8
func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}

}

// task-9
func reverse(r int) {
	temp := r
	for r != 0 {
		temp = r % 10
		fmt.Print(temp)
		r /= 10
	}
	fmt.Println()
}
