package main

import "fmt"

// task-1
func tubSon(t int) {
	counter := 0
	for j := 2; j <= t/2; j++ {
		if t%j == 0 {
			counter++
		}
	}
	if counter == 0 {
		fmt.Println(t, " tub son.")
	} else {
		fmt.Println(t, " tub son emas.")
	}
}

// task-2
func swap(N int) {
	// temp := 0
	var temp int

	for N != 0 {
		temp = N % 10
		fmt.Printf("%d ", temp)
		N /= 10
	}
}

// task-3
func sumNum(N int) {
	var count = 0
	tempN := 0
	tempN = N
	for tempN != 0 {
		tempN /= 10
		count++
	}
	temp := 1
	sum := 0
	for i := 0; i < count; i++ {
		if i == 0 {
			for j := 1; j < count; j++ {
				temp *= 10
			}
		}
		sum += N / temp
		N %= temp
		temp /= 10
	}
	fmt.Println(sum)

}

// task-4
func sumN(n int) {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += i
		fmt.Print(i, " + ")
	}

	fmt.Println(" = ", sum)
}

// task-5
func print(a, b int) {
	if a < b {
		for i := a; i <= b; i++ {
			fmt.Print(i, " ")
		}
	} else if a > b {
		for j := a; j >= b; j-- {
			fmt.Print(j, " \n")
		}
	}
	fmt.Println()
}

// task-6
func evenNum(n int) {
	for i := n; i > 1; i-- {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// task-7
func middleNum(a, b, c int) {
	if a > c && b < c {
		fmt.Println(c)
	} else if a > b && c < b {
		fmt.Println(b)
	} else if b > a && c < a {
		fmt.Println(a)
	} else if b > c && a < c {
		fmt.Println(c)
	} else if c > b && a < b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}

}

// task-8
func evenSum(n int) bool {
	tempN := 0
	tempN = n
	count := 0
	for tempN != 0 {
		tempN /= 10
		count++
	}
	var temp = 1
	var sum = 0
	for i := 0; i < count; i++ {
		if i == 0 {
			for j := 1; j < count; j++ {
				temp *= 10
			}
		}
		sum += n / temp

		n %= temp
		temp /= 10
	}
	if sum%2 == 0 {
		return true
	} else {
		return false
	}

}

// task-9
func divideN(n, n1 int) int {
	return n / n1
}

// task-10
func sumNums(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += i
	}

	return sum
}

// task-11
func algorithm(v1 int) {
	temp := 0
	temp = v1
	count := 0
	for temp != 0 {
		temp = temp / 10
		count++
	}

	tempN := 1
	for i := 0; i < count; i++ {
		if i == 0 {
			for j := 1; j < count; j++ {
				tempN *= 10
			}
		}
		fmt.Print(v1/tempN, " ")

		v1 %= tempN
		tempN /= 10
	}
	fmt.Println()
}

// task-12
func primeDivider(n int) {
	for i := 2; i <= n; i++ {
		var sum int = 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				sum++
			}
		}
		if sum == 0 {
			if n%i == 0 {
				fmt.Print(i, " ")
			}
		}
	}
	fmt.Println()
}

// task-13
func multiplication(n int) int {
	m := 1
	for i := 1; i <= n; i++ {
		m *= i
	}
	return m
}

// task-14
func add(a, b, c int) int {
	if a+b == c {
		return 1
	} else {
		return 0
	}
}

// task-15
func power(n float32, m int) float32 {
	for i := 1; i < m; i++ {
		n *= n
	}
	return n
}

// task-16
func toqSon(n int) {
	for i := 1; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// task-17
func multipTable(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
}
func main() {
	//task-1
	var tub int
	fmt.Print(("Input number: "))
	fmt.Scan(&tub)
	tubSon(tub)

	//task-2
	var N int
	fmt.Print(("Input number: "))
	fmt.Scan(&N)
	swap(N)

	//task-3
	var N1 int
	fmt.Print(("Input number: "))
	fmt.Scan(&N1)
	sumNum(N1)

	//task-4
	var n int
	fmt.Print("Input num: ")
	fmt.Scan(&n)
	sumN(n)

	//task-5
	var a, b int
	fmt.Print("Input numbers: ")
	fmt.Scan(&a, &b)
	print(a, b)

	// //task-6
	var n1 int
	fmt.Print("Input number: ")
	fmt.Scan(&n1)
	evenNum(n1)

	//task-7
	var a1, b1, c int
	fmt.Print("Input three numbers: ")
	fmt.Scan(&a1, &b1, &c)
	middleNum(a1, b1, c)

	//task-8
	var g int
	fmt.Print("Input num: ")
	fmt.Scan(&g)
	fmt.Println(evenSum(g))

	//task-9
	var n2, n3 int
	fmt.Print("Input numbers: ")
	fmt.Scan(&n2, &n3)
	fmt.Println(divideN(n2, n3))

	//task-10
	var v int
	fmt.Print("Input num: ")
	fmt.Scan(&v)
	fmt.Println(sumNums(v))

	//task-11
	var v1 int
	fmt.Print("Input num: ")
	fmt.Scan(&v1)
	algorithm(v1)

	//task-12
	var n4 int
	fmt.Print("Input num: ")
	fmt.Scan(&n4)
	primeDivider(n4)

	//task-13
	var m int
	fmt.Print("Input num: ")
	fmt.Scan(&m)
	fmt.Println(multiplication(m))

	//task-14
	var k, k1, k2 int
	fmt.Print("Input three numbers: ")
	fmt.Scan(&k, &k1, &k2)
	fmt.Println(add(k, k1, k2))

	//task-15
	var r float32
	var r1 int
	fmt.Print("Input numbers: ")
	fmt.Scan(&r, &r1)
	fmt.Println(power(r, r1))

	//task-16
	var h int
	fmt.Print("Input number: ")
	fmt.Scan(&h)
	toqSon(h)

	//task-17
	var table int
	fmt.Print("Input number: ")
	fmt.Scan(&table)
	multipTable(table)
}
