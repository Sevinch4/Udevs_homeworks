package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error when opening file:", err.Error())
		return
	}
	defer file.Close()
	f, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error when read file:", err.Error())
		return
	}
	nums := make([]int, 0)
	words := make([]string, 0)
	slice := strings.Split(string(f), " ")//_ bolib slice iciga saqlab oladi
	for _, v := range slice {//slice boica aylanadi
		a, err := strconv.Atoi(v)//stringni int ga convert qiladi
		if err != nil {//qila olmasa error otadi, chunki son emas harf ishtirok etgan bolishi mm
			words = append(words, v)//words slicega saqlavoladi vni qiymatini
		} else {//convert qila olsa nums slicega saqlab oladi
			nums = append(nums, a)
		}
	}
	fmt.Println( words, nums)
}
