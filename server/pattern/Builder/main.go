package main

import "fmt"

// What is Builder Pattern?
// Builder pattern is a pattern of creational design pattern which helps us to build an complex object step by step.
// It means we can set our propeties when we actually need them in our application code.

// Why:
// Some objects are easy to create with simple steps but some need special ceremonies to construct an object.
// This is useful when you dont want to serve the unncessory information to your end user.

// When do we need this?
// When you don't need your end-user to provide millions of details in parameters of a constructor and you want to make his/her life easy by providing them with a simple way to create a complex object.

type Person struct {
	name, address, pin             string
	workaddress, company, position string
	salary                         int
}

type PersonBuilder struct {
	Person *Person
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{Person: &Person{}}
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func (a *PersonAddressBuilder) At(address string) *PersonAddressBuilder {
	a.Person.address = address
	return a
}

func (a *PersonAddressBuilder) WithPostalCode(code string) *PersonAddressBuilder {
	a.Person.pin = code
	return a
}

func (j *PersonJobBuilder) Name(name string) *PersonJobBuilder {
	j.Person.name = name
	return j
}

func (j *PersonJobBuilder) As(position string) *PersonJobBuilder {
	j.Person.position = position
	return j
}

func (j *PersonJobBuilder) For(company string) *PersonJobBuilder {
	j.Person.company = company
	return j
}

func (j *PersonJobBuilder) In(companyAddress string) *PersonJobBuilder {
	j.Person.workaddress = companyAddress
	return j
}

func (j *PersonJobBuilder) Salary(salary int) *PersonJobBuilder {
	j.Person.salary = salary
	return j
}

func (b *PersonBuilder) Build() *Person {
	return b.Person
}

func main() {
	pb := NewPersonBuilder()

	pb.Lives().
		At("314 Ave Seattle").
		WithPostalCode("234567").
		Works().
		Name("John Smith").
		As("Software Engineer").
		For("Microsoft").
		In("Banglore").
		Salary(120000)

	person := pb.Build()
	fmt.Printf("%v \n", person)
}
