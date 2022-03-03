package main

import (
    "fmt"
    "time"
	"rsol.com/hello/articles"
)

import "rsc.io/quote"


type Meetup struct {
	location string
	city     string
	date     time.Time
	people   []People
}

func (m Meetup) MeetupPeople() {
	for _, v := range m.people {
		v.SayHello()
		v.GetDetails()
	}
}

func main() {
	fmt.Println(quote.Go())
	fmt.Print(quote.Glass())

	p := Person{
		name:  "shiju",
		age:   35,
		city:  "Kochi",
		phone: "+900sd0sad",
	}

    p1 := Person{
		name:  "sam",
		age:   20,
		city:  "TX",
		phone: "+000000000",
	}

	m := Meetup{
        location: "",
        city: "",
        date: time.Now(),
        people: []People {p,p1},
    }


    m.MeetupPeople()

	fmt.Println("..." + art.Name)
}
