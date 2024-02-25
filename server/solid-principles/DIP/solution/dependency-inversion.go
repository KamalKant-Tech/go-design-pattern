package main

import "fmt"

// This demonstrates the Dependency Inversion Principle, as the Department struct depends on an abstraction (the Employee interface),
// rather than on a specific implementation (the Worker or Supervisor struct).
// This makes the code more flexible and easier to maintain, as changes to the implementation of workers and supervisors will not affect the Department struct.

type Employee interface {
	GetID() int
	GetName() string
}

type Worker struct {
	Id   int
	Name string
}

type Supervisor struct {
	Id   int
	Name string
}

type Department struct {
	Employee []Employee
}

func (w *Worker) GetID() int {
	return w.Id
}

func (w *Worker) GetName() string {
	return w.Name
}

func (s *Supervisor) GetID() int {
	return s.Id
}

func (s *Supervisor) GetName() string {
	return s.Name
}

func (d *Department) GetEmployeeNames() (res []string) {
	for _, e := range d.Employee {
		res = append(res, e.GetName())
	}
	return res
}

func (d *Department) GetEmployee(id int) Employee {
	for _, e := range d.Employee {
		if e.GetID() == id {
			return e
		}
	}
	return nil
}

func main() {
	w := Worker{1, "John"}
	s := Supervisor{2, "Milley"}

	dept := Department{[]Employee{&w, &s}}
	fmt.Println(dept.GetEmployeeNames())
	emp := dept.GetEmployee(1)

	switch v := emp.(type) {
	case *Worker:
		fmt.Printf("Found Worker employee id %v", v.Id)
	case *Supervisor:
		fmt.Printf("Found Supervisor employee id %v", v)
	default:
		fmt.Println("No employee Found")
	}
}
