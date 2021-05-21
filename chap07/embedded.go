package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployee() []Employee {
	// do some buisness logic
	return m.Reports
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	fmt.Println(m)
	fmt.Println("ID:", m.ID)
	fmt.Println("Name:", m.Name)
	fmt.Println("Reports:", m.Reports)
	fmt.Println(m.Description())
}
