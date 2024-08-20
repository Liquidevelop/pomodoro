package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

var timeRemaining time.Duration

func startTimer(duration time.Duration, done chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	timeRemaining = duration

	for {
		select {
		case <-ticker.C:
			timeRemaining -= time.Second
			if timeRemaining <= 0 {
				done <- true
				return
			}
		}
	}
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func main() {
	input := "short"

	fmt.Print("Enter pomodoro type short/long (short): ")
	fmt.Scanln(&input)

	var duration time.Duration

	switch input {
	case "long":
		duration = 25 * time.Minute
	case "short":
		duration = 5 * time.Minute
	default:
		fmt.Println("Invalid input, please enter short or long")
	}

	done := make(chan bool)
	go startTimer(duration, done)

	fmt.Printf("Starting a %v Pomodoro...\n", formatDuration(duration))

	for {
		fmt.Print("Enter 'status' to check remaining time or 'exit' to quit: ")
		fmt.Scanln(&input)

		switch input {
		case "status":
			fmt.Printf("Time remaining: %v\n", formatDuration(timeRemaining))
		case "exit":
			fmt.Println("Exiting the Pomodoro timer.")
			return
		default:
			fmt.Println("Unknown command, please enter 'status' or 'exit'.")
		}

		select {
		case <-done:
			fmt.Println("Pomodoro completed!")
			beeep.Notify("Pomodoro completed!", "Time to take a break", "")
			return
		default:

		}
	}
}
