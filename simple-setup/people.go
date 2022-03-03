package main

import (
    "fmt"
)

type Person struct {
	name        string
	age         int
	city, phone string
}

type People interface {
	SayHello()
	GetDetails()
}

func (p Person) SayHello() {
	fmt.Printf("\nHi, I am %s, from %s\n", p.name, p.city)
}

func (p Person) GetDetails() {
	fmt.Printf("[Name: %s, Age: %d, City: %s, Phone: %s]\n", p.name, p.age, p.city, p.phone)
}