package main

import "fmt"

// What is DIP(Dependency inversion Principle)

// Dependency inversion state that high lever module should not depend on the low level module, but rather both should depend on the abstractions.

// Why DIP: This hepls to recude the coupling b/w the modules and make the code more flexible and maintainable.

// Why Not DIP:
// 1. If HLM depends on the low level modules there high chances of code break in HLM modules which cause the crash of the application.
// 2. Its hard to write the testing for the high level modules and component of HLM can not be replaceble by Mock LLM modules.

// How To implement DIP:
// Create an interface for the LLM and inject the type of interface into high level modules. Which can be replaceble by the type which implements that interface.

// Low Level Module
type worker struct {
	Id   int
	Name string
}

// Low Level Module
type supervisor struct {
	Id   int
	Name string
}

// High level module
// This below struct violating the rule of DIP cause Department is directly depend on the low level modules.
type Department struct {
	Worker     worker
	Supervisor supervisor
}

func (w *worker) GetID() int {
	return w.Id
}

func (w *worker) GetName() string {
	return w.Name
}

func (s *supervisor) GetId() int {
	return s.Id
}

func (s *supervisor) GetName() string {
	return s.Name
}

func main() {
	w := worker{1, "John"}
	s := supervisor{2, "Milley"}
	dept := &Department{w, s}

	fmt.Println(dept.Worker.GetID())
}
