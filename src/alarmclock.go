package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) func() {
	for {
		fmt.Println("The time is " + time.Now().Format("15:04:05") + ": " + text)
		time.Sleep(delay)
	}
}

func main() {
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	Remind("Time to sleep", 60*time.Second)
}
