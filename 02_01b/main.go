package main

import (
	"flag"
	"log"
	"sync"
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
	var print1 = func(message string, wg *sync.WaitGroup) {
		defer wg.Done()
		println(message)
	}
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go print1(message, &wg)
	}
	wg.Wait()
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		log.Println(m)
		repeat(int(*factor), m)
	}
}
