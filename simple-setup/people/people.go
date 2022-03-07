package people

import (
    "fmt"
)

type Person struct {
	Name        string
	Age         int
	City, Phone string
}

type People interface {
	SayHello()
	GetDetails()
}

func (p Person) SayHello() {
	fmt.Printf("\nHi, I am %s, from %s\n", p.Name, p.City)
}

func (p Person) GetDetails() {
	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", p.Name, p.Age, p.City, p.Phone)
}