package main

import "fmt"

func main() {
	//if

	//task-1
	fmt.Println("Task-1")
	var (
		A int
		B int
	)
	fmt.Print("Input numbers: ")
	fmt.Scan(&A, &B)
	//A son kichik B son katta blsh kere
	if A > B {
		A, B = B, A
		fmt.Println(A, B)
	} else {
		fmt.Println(A, B)
	}

	//task-2
	fmt.Println("Task-2")
	var (
		a int
		b int
	)
	fmt.Print("Input numbers: ")
	fmt.Scan(&a, &b)
	if a != b {
		a, b = a+b, a+b
		//b = b + a
		fmt.Println(a, b)
	} else {
		a = 0
		b = 0
		fmt.Println(a, b)
	}

	//task-3
	fmt.Println("Task-3")
	var (
		a1 int
		b1 int
	)
	fmt.Print("Input numbers: ")
	fmt.Scan(&a1, &b1)
	if a1 < b1 {
		a1 = b1
		fmt.Println(a1, b1)
	} else if a1 > b1 {
		b1 = a1
		fmt.Println(a1, b1)
	} else {
		a1 = 0
		b1 = 0
		fmt.Println(a1, b1)
	}

	//task-4
	fmt.Println("Task-4")
	var (
		n1 int
		n2 int
		n3 int
	)
	fmt.Print("Input numbers: ")
	fmt.Scan(&n1, &n2, &n3)
	if n1 > n2 && n1 > n3 {
		fmt.Println(n1)
	} else if n2 > n1 && n2 > n3 {
		fmt.Println(n2)
	} else {
		fmt.Println(n3)
	}

	//task-5
	fmt.Println("Task-5")
	var (
		k1 int
		k2 int
		k3 int
	)
	fmt.Print("Input numbers: ")
	fmt.Scan(&k1, &k2, &k3)
	if (k1 > k2 && k1 < k3) || (k1 < k2 && k1 > k3) {
		fmt.Println(k1)
	} else if (k2 > k1 && k2 < k3) || (k2 < k1 && k2 > k3) {
		fmt.Println(k2)
	} else {
		fmt.Println(k3)
	}

	//task-6
	fmt.Println("Task-6")
	var (
		son1 int
		son2 int
		son3 int
	)
	fmt.Print("Input three numbers: ")
	fmt.Scan(&son1, &son2, &son3)
	if son1 > son2 && son1 > son3 && son2 > son3 {
		fmt.Println(son3, son1)
	} else if son1 > son2 && son1 > son3 && son2 < son3 {
		fmt.Println(son2, son1)
	} else if son2 > son1 && son2 > son3 && son1 > son3 {
		fmt.Println(son3, son2)
	} else if son2 > son1 && son2 > son3 && son1 < son3 {
		fmt.Println(son1, son2)
	} else if son3 > son1 && son3 > son2 && son1 > son2 {
		fmt.Println(son2, son3)
	} else {
		fmt.Println(son1, son3)
	}

	//task-7
	fmt.Println("Task-7")
	var a, b, c int
    fmt.Print("Input three numbers: ")
    fmt.Scan(&a,&b,&c)
    sum1 := a + b
    sum2 := b + c
    sum3 := a + c
    maxSum := sum1
    if sum2 > maxSum {
        maxSum = sum2
    }
    if sum3 > maxSum {
        maxSum = sum3
	}
    fmt.Printf("Max sum : %d\n", maxSum)
}

	//task-8
	fmt.Println("Task-8")
	var A1, B1, C float64
    fmt.Print("Input three numbers: ")
    fmt.Scan(&A1,&B1,&C)
    if A1 <= B1 && B1 <= C {
        fmt.Printf("O'sish tartibida: A = %v, B = %v, C = %v\n", A1, B1, C)
    } else {
        if A1 > B {
            A1, B1 = B1, A1
        }
        if B1 > C {
            B1, C = C, B1
        }
        if A1 > B1 {
            A1, B1 = B1, A1
        }
        fmt.Printf("Ikkinchilantirilgan tartibida: %v, %v, %v\n", A1, B1, C1)
    }
	//task-9
	fmt.Println("Task-9")
	var (
		A2 int
		B2 int
		C2 int
	)
	fmt.Print("Input three num : ")
	fmt.Scan(&A2,&B2,&C2)
	if A2 <= B2 && B2 <= C2 {
        fmt.Printf("O'sish tartibida: A = %v, B = %v, C = %v\n", A2, B, C2)
    } else if A2 >= B2 && B2 >= C2 {
        fmt.Printf("Kamayish tartibida: A = %v, B = %v, C = %v\n", A2, B2, C2)
    } else {
        if A2 > B2 {
            A2, B2 = B2, A2
        }
        if B2 > C2 {
            B2, C2 = C2, B2
        }
        if A2 > B2 {
            A2, B2 = B2, A2
        }
        fmt.Printf("Ikkilantirilgan tartibida: %v, %v, %v\n", A2, B2, C2)
    }
	fmt.Println("Task-10")
	var (
		l int
		l1 int
		l2 int
	)
	fmt.Print("Input three num : ")
	fmt.Scan(&l,&l1,&l2)
	if l == l1 {
		fmt.Printf("Tartib raqami: %d %d %d\n", l, l1, l2)
	} else if l == l2 {
		fmt.Printf("Tartib raqami: %d %d %d\n", l, l2, l1)
	} else if l1 == l2 {
		fmt.Printf("Tartib raqami: %d %d %d\n", l1, l2, l1)
	} else {
		if l > l1 && l > l2 {
			fmt.Printf("Tartib raqami: %d %d %d\n", l1, l2, l)
		} else if l1 > l && l1 > l2 {
			fmt.Printf("Tartib raqami: %d %d %d\n", l, l2, l1)
		} else {
			fmt.Printf("Tartib raqami: %d %d %d\n", l, l1, l2)
		}
	}
	fmt.Println("Task-11")
	var(
		a2 int
		b2 int 
		c2 int
		d int
	)
	fmt.Print("Input four num : ")
	fmt.Scan(&a2,&b2,&c2,&d)
	if a2 == b2 && b2 == c2 {
		fmt.Printf("Tartib raqami: %d %d %d %d\n", a2, b2, c2, d)
	} else if a2 == b2 && b2 == d {
		fmt.Printf("Tartib raqami: %d %d %d %d\n", a2, b2, d, c2)
	} else if a2 == c2 && c2 == d2 {
		fmt.Printf("Tartib raqami: %d %d %d %d\n", a2, c2, d, b2)
	} else if b2 == c2 && c2 == d2 {
		fmt.Printf("Tartib raqami: %d %d %d %d\n", b2, c2, d, a2)
	} else {
		if a2 == b2 {
			fmt.Printf("Tartib raqami: %d %d %d %d\n", a2, b2, d, c2)
		} else if a2 == c2 {
			fmt.Printf("Tartib raqami: %d %d %d %d\n", a2, c2, d, b2)
		} else if a2 == d {
			fmt.Printf("Tartib raqami: %d %d %d %d\n", a2, d, b2, c2)
		} else if b2 == c2 {
			fmt.Printf("Tartib raqami: %d %d %d %d\n", b2, c2, d, a2)
		} else if b2 == d {
			fmt.Printf("Tartib raqami: %d %d %d %d\n", b2, d, a2, c2)
		} else {
			fmt.Printf("Tartib raqami: %d %d %d %d\n", c2, d, a2, b2)
		}
	}

	//for

	//task - 1
	fmt.Println("Task-1")
	var n int
	fmt.Print("Input n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}

	//task-2
	fmt.Println("Task-2")
	var N int
	fmt.Print("Input n: ")
	fmt.Scan(&N)
	for i := 1; i <= N; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}

	//task-3
	fmt.Println("Task-3")
	fmt.Print("Tub sonlar: ")
	for i := 10; i <= 99; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}

		}
		if isPrime {
			fmt.Print(i, " ")
		}

	}

	//task-4
	fmt.Println("Task-4")
	var (
		yuz int
		on  int
		bir int
	)
	for i := 100; i <= 999; i++ {
		yuz = i / 100
		on = (i / 10) % 10
		bir = i % 10
		if yuz == on || yuz == bir || bir == on {
			fmt.Print(i, " ")
		}
	}

	//task-5????????????????
	fmt.Println("Task-5")
	var (
		num1 int
	)
	fmt.Print("Input num: ")
	fmt.Scan(&num1)
	temp := num1
	sum := 0
	for temp > 0 {
		temp /= 10
		sum++
	}
	fmt.Println(sum)
	var i = 1
	for i = 1; i <= sum; i *= 10 {
		i *= 10
	}
	fmt.Println(i)
	for num1 >= 10 {
		num1 /= i
		fmt.Print(num1, " ")
	}

	//fmt.Println()

	// //task-6 ?????????????
	fmt.Println("Task-6")
	//var i int
	for i := 1; i <= 6; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")

		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//task-7 ??
	fmt.Println("Task-7")
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if i == 1 || i == 4 || j == 1 || j == 4 {
				fmt.Print("* ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//switch

	//task-1
	var (
		Y string
		K int
	)
	fmt.Print("Input y and k: ")
	fmt.Scan(&Y, &K)
	switch Y {
	case "s":
		switch K {
		case 0:
			fmt.Println(Y)
		case 1:
			fmt.Println("q-sharq")
		case 2:
			fmt.Println("g-garb")
		}
	case "j":
		switch K {
		case 0:
			fmt.Println(Y)
		case 1:
			fmt.Println("g-garb")
		case 2:
			fmt.Println("q-sharq")
		}
	case "q":
		switch K {
		case 0:
			fmt.Println(Y)
		case 1:
			fmt.Println("j-janub")
		case 2:
			fmt.Println("s-shimol")
		}
	case "g":
		switch K {
		case 0:
			fmt.Println(Y)
		case 1:
			fmt.Println("s-shimol")
		case 2:
			fmt.Println("j-janub")
		}

	}
	//task-2
	var (
		c  string
		K1 string
		K2 int
	)
	fmt.Print("Input locaters location: ")
	fmt.Scan(&c)
	fmt.Println()
	fmt.Print("Input commands: ")
	fmt.Scan(&K1, &K2)
	switch c {
	case "s":
		switch K1 {
		case "s":
			//shimol
			switch K2 {
			case 0:
				c = "g-garb"
				//onga buril
			case 1:
				c = "q-sharq"
				//chapga buril
			case 2:
				c = "j-janub"
				//180 burilish

			}
		case "j":
			//janub
			switch K2 {
			case 0:
				c = "q-sharq"
				//onga buril
			case 1:
				c = "g-g'arb"
				//chapga buril
			case 2:
				c = "s-shimol"
				//180 burilish
			}
		case "q":
			//sharq
			switch K2 {
			case 0:
				c = "s-shimol"
				//onga buril
			case 1:
				c = "j-janbu"
				//chapga buril
			case 2:
				c = "g-g'arb"
				//180 burilish
			}
		case "g":
			//g'arb
			switch K2 {
			case 0:
				c = "g-garb"
				//onga buril
			case 1:
				c = "q-sharq"
				//chapga buril
			case 2:
				c = "j-janub"
				//180 burilish
			}
		}

	case "j":
	case "q":
	case "g":

	}
