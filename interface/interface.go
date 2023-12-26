package main

import "fmt"

func main(){
	p := Programmer{
		JobName: "Programmer",
	}
	t := Teacher{
		JobName: "Teacher",
	}
	Work(p)
	Work(t)				
}

type Worker interface{
	Work()
}

func Work(w Worker){
	w.Work()
}

func (p Programmer) Work(){
	fmt.Println("Job - ",p.JobName)
}

func (t Teacher) Work(){
	fmt.Println("Job - ",t.JobName)
}

type Programmer struct{
	JobName string
}

type Teacher struct{
	JobName string
}
