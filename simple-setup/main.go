package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
	"rsol.com/simple-setup/articles"
	"rsol.com/simple-setup/people"
)

type Meetup struct {
	location string
	city     string
	date     time.Time
	people   []people.People
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

	p := people.Person{
		Name:  "shiju",
		Age:   35,
		City:  "Kochi",
		Phone: "+900sd0sad",
	}

    p1 := people.Person{
		Name:  "sam",
		Age:   20,
		City:  "TX",
		Phone: "+000000000",
	}

	m := Meetup{
        location: "",
        city: "",
        date: time.Now(),
        people: []people.People {p,p1},
    }


    m.MeetupPeople()

	fmt.Println("..." + art.Name)
}
