package main

import "fmt"

type Employee struct {
	Designation string
	Salary int
}

func updateSalary(e Employee, c chan Employee) {
	switch e.Designation {
    case "AP":
        e.Salary += 3000
    case "AP(Sr.Gr.)":
        e.Salary += 5000
    case "AP(SG)":
        e.Salary += 7000
    case "AsP":
        e.Salary += 7500
    case "Prof":
        e.Salary += 6000
    }

	c <- e
}

func main() {
	employees_list := []Employee{
		{"AP", 100000},
		{"AP(Sr.Gr)", 150000},
		{"Prof", 200000},
	} 

	c := make(chan Employee)

	for _, emp := range employees_list {
		go updateSalary(emp,c)
	}

	for i := 0; i < len(employees_list); i++ {
        e := <-c
        fmt.Printf("Designation: %s, Updated Salary: %d\n", e.Designation, e.Salary)
    }
}