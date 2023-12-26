package main

import (
	"fmt"
)

//task-1
type Student struct{
	name string
	scholarship int
	course int
}

//task-2
type Student2 struct{
	name string 
	scholar int 
	course int
}
//task-3
type AEROPORT struct{
	planeType string 
	flightDate string 
	fromCity string
	toCity string 
	flightTime int
}
//task-5
type Team struct{
	name string
	coach string
	playersCount int
	Captain string 
}
func main(){
	fmt.Println("Task-1")
	student := []Student{
		{
			name: "Sevinch",
			scholarship: 300,
			course: 2,
		},
		{
			name: "Sarvinoz",
			scholarship: 500,
			course: 2,
		},
		{
			name: "mushtariy",
			scholarship: 300,
			course: 1,
		},
		{
			name: "azam",
			scholarship: 700,
			course: 2,
		},
		{
			name: "munisa",
			scholarship: 0,
			course: 1,
		},
		{
			name: "Sardor",
			scholarship: 900,
			course: 2,
		},
		{
			name: "Kamron",
			scholarship: 200,
			course: 2,
		},
		{
			name: "Svetlana",
			scholarship: 400,
			course: 1,
		},
		{
			name: "Mirshod",
			scholarship: 300,
			course: 2,
		},
		{
			name: "Sarvar",
			scholarship: 900,
			course: 2,
		},}
	sum := 0
	for _,c :=range student{
		if c.course == 2{
			sum += c.scholarship
		}
	}
	fmt.Println("2-course scholar is : ",sum)

	//task-2
	fmt.Println("Task-2")
	student1 := []Student2{
		{
			name: "Sevinch",
			scholar: 300,
			course: 2,
		},
		{
			name: "Aisha",
			scholar: 500,
			course: 2,
		},
		{
			name: "Asli",
			scholar: 300,
			course: 1,
		},
		{
			name: "azam",
			scholar: 700,
			course: 2,
		},
		{
			name: "munisa",
			scholar: 0,
			course: 1,
		},
		{
			name: "Sardor",
			scholar: 900,
			course: 2,
		},
		{
			name: "Ali",
			scholar: 200,
			course: 2,
		},
		{
			name: "Ece",
			scholar: 400,
			course: 1,
		},
		{
			name: "Mirshod",
			scholar: 300,
			course: 2,
		},
		{
			name: "Sarvar",
			scholar: 900,
			course: 2,
		},
	}
	for _,s := range student1{
		if len(s.name) < 5{
			fmt.Printf("Name: %s  Course: %d  Scholar: %d\n",s.name,s.course,s.scholar)
		}
	}
	//task-3
	aeroport := []AEROPORT{
		{
			planeType: "Aircraft type",
			flightDate: "June",
			fromCity: "Tashkent",
			toCity: "Samarkand",
			flightTime: 4,
		},
		{
			planeType: "Geometry",
			flightDate: "July",
			fromCity: "Samarqand",
			toCity: "Xiva",
			flightTime: 5,
		},
		{
			planeType: "Woodworking",
			flightDate: "Avgust",
			fromCity: "Tashkent",
			toCity: "Saudia Arabiya",
			flightTime: 9,
		},
		{
			planeType: "Aircraft type",
			flightDate: "December",
			fromCity: "Tashkent",
			toCity: "USA",
			flightTime: 10,
		},
		{
			planeType: "Woodworking",
			flightDate: "September",
			fromCity: "Italy",
			toCity: "Antalya",
			flightTime: 8,
		},
		{
			planeType: "Geometry",
			flightDate: "July",
			fromCity: "Koreya",
			toCity: "Tashkent",
			flightTime: 2,
		},
		{
			planeType: "Aircraft type",
			flightDate: "June",
			fromCity: "German",
			toCity: "Tashkent",
			flightTime: 3,
		},
		{
			planeType: "Woodworking",
			flightDate: "October",
			fromCity: "Monako",
			toCity: "Tashkent",
			flightTime: 2,
		},
		{
			planeType: "Aircraft type",
			flightDate: "November",
			fromCity: "New york",
			toCity: "Samarkand",
			flightTime: 10,
		},
		{
			planeType: "Aircraft type",
			flightDate: "November",
			fromCity: "New zealand",
			toCity: "Tashkent",
			flightTime: 3,
		},

	}
	fmt.Println("In the summer:")
	for _,a := range aeroport{
		if "July" == a.flightDate{
			fmt.Printf("Plane type: %s\nFlight date: %s\nFrom city: %s\nTo city: %s\nFlight time: %d\n",a.planeType,a.flightDate,a.fromCity,a.toCity,a.flightTime)
		}else if "June" == a.flightDate{
			fmt.Printf("Plane type: %s\nFlight date:%s\nFrom city: %s\nTo city: %s\nFlight time: %d\n",a.planeType,a.flightDate,a.fromCity,a.toCity,a.flightTime)
		}else if "Avgust" == a.flightDate{
			fmt.Printf("Plane type: %s\nFlight date: %s\nFrom city: %s\nTo city: %s\nFlight time: %d\n",a.planeType,a.flightDate,a.fromCity,a.toCity,a.flightTime)
		}
	}
	fmt.Println("Arrive to Tashkent:")
	//task-4
	for _,a1 := range aeroport{
		if "Tashkent" == a1.toCity && (a1.flightTime == 2 || a1.flightTime == 3){
			fmt.Printf("Plane type: %s\nFlight date: %s\nFrom city: %s\nTo city: %s\nFlight time: %d\n",a1.planeType,a1.flightDate,a1.fromCity,a1.toCity,a1.flightTime)
		}
	}
	//task-5
	team := []Team {
		{
			name: "aircraft",
			coach: "sport",
			playersCount: 28,
			Captain: "someone",
		},
		{
			name: "bears",
			coach: "sport",
			playersCount: 20,
			Captain: "someone",
		},
		{
			name: "cars",
			coach: "transportation",
			playersCount: 10,
			Captain: "someone",
		},
		{
			name: "craft",
			coach: "sport",
			playersCount: 15,
			Captain: "someone",
		},
		{
			name: "pull",
			coach: "something",
			playersCount: 50,
			Captain: "military",
		},
	}
	var name string
	var n int
	var slice []string
	slice1 := []Team{}
	fmt.Print("How many teams you want input (1-5) : ")
	fmt.Scan(&n)
	fmt.Print("Input team name: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		slice = append(slice, name)
	}
	for _,t := range team{
		for i := 0; i < len(slice); i++ {
			if slice[i] == t.name{
				slice1 = append(slice1, t)
			}
		}
	}
	
	for i:=0;i<len(slice1)-1;i++ {
		for j := i+1; j < len(slice1); j++ {
			if 	slice1[i].playersCount<slice1[j].playersCount{
				slice1[j],slice1[i]=slice1[i],slice1[j]
			}
		}
	}
	for _,val := range slice1{
		fmt.Println(val)
	}


}