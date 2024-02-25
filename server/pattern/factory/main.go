package main

import "fmt"

type Product interface {
	PrintName() string
}

type Electronics struct {
	Name string
}

type Furniture struct {
	Name string
}

// lets implement the product interface for Electronics
func (e *Electronics) PrintName() string {
	return e.Name
}

// Lets implement the product interface for furniture
func (f *Furniture) PrintName() string {
	return f.Name
}

func CreateProduct(category string) Product {
	switch category {
	case "E":
		return &Electronics{Name: "Mobile"}
	case "F":
		return &Furniture{Name: "Chair"}
	default:
		return nil
	}
}
func main() {
	e := CreateProduct("E")
	fmt.Println(e.PrintName())

	f := CreateProduct("F")
	fmt.Println(f.PrintName())
}
