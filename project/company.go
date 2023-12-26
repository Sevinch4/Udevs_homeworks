package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Company struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Owner   Owner    `json:"owner"`
	Workers []Worker `json:"workers"`
}

type Owner struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Worker struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Salary    int    `json:"salary"`
	Level     string `json:"level"`
}
var companies1 = []Company{
	{
		Id: 1,
		Name: "apple",
		Owner: Owner{
			FirstName: "steve",
			LastName: "jobs",
		},
		Workers: []Worker{
			{
				ID: 1,
				FirstName: "somename",
				LastName: "somelastname",
				Salary: 30000,
			},
		},
	},
}

func main() {
	companies := []Company{}
	file, err := os.OpenFile("test.json", os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&companies); err != nil {
		panic(err)
	}
	//	fmt.Println(companies)
	//task-1
	sort.Slice(companies, func(i, j int) bool {
		return len(companies[i].Workers) > len(companies[j].Workers)
	})
	//task-2
	slice := make([]Worker, 0)
	for _, v := range companies {
		for _, wor := range v.Workers {
			slice = append(slice, wor)
		}
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Salary > slice[j].Salary
	})
	//task-3
	sort.Slice(companies,func(i, j int) bool {
		sumI,sumJ := 0,0
		for _,worker := range companies[i].Workers{
			sumI += worker.Salary
		}
		fmt.Println(sumI,companies[i].Name)
		for _,worker := range companies[j].Workers{
			sumJ += worker.Salary
		}
		fmt.Println(sumJ,companies[j].Name)
		return sumI > sumJ

	})
	for _,v := range companies{
		fmt.Println(v.Owner)
	}
	//task-1
	// for _,company := range companies{
	// 	fmt.Println(company.Name)
	// }

	//task-2
	// for i := 0; i <= 3; i++ {
	// 	fmt.Println(slice[i])
	// }

}
