package main

import (
	"flag"
	"log"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string) {
	print1 := func(message string, ch chan struct{}) {
		println(message)
		ch <- struct{}{}
	}
	ch := make(chan struct{})
	for i := 0; i < n; i++ {
		go print1(message, ch)
	}
	for i := 0; i < n; i++ {
		<-ch
	}

}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		log.Println(m)
		repeat(int(*factor), m)
	}
}
