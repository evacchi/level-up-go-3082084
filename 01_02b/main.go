package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	lines := strings.Split(msg, " ")
	for _, l := range lines {
		acc := []rune{}
		for i, c := range l {
			for j := 0; j <= i; j++ {
				acc = append(acc, c)
			}
		}
		print(string(acc))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
