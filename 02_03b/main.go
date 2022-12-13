package main

import (
	"fmt"
	"log"
)

// the number of attendees we need to serve lunch to
const consumerCount = 300

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string, course []chan string, signal chan struct{}) {
	for _, ch := range course {
		log.Printf("%s eats %s.\n", name, <-ch)
	}
	signal <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string, outCourse chan string, signal chan struct{}) {
	for {
		select {
		case outCourse <- course:
		case <-signal:
			return
		}
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)
	var courses []chan string
	doneEating := make(chan struct{})
	doneServing := make(chan struct{})
	for _, c := range foodCourses {
		ch := make(chan string)
		courses = append(courses, ch)
		go serveLunch(c, ch, doneServing)
	}
	for i := 0; i < consumerCount; i++ {
		name := fmt.Sprintf("Attendee %d", i)
		go takeLunch(name, courses, doneEating)
	}
	for i := 0; i < consumerCount; i++ {
		<-doneEating
	}
	close(doneServing)

}
