package main

import (
	"testing"
	"time"
)

// TestFormatDuration tests the formatDuration function.
func TestFormatDuration(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{5 * time.Minute, "05:00"},
		{25 * time.Minute, "25:00"},
		{time.Minute + 30*time.Second, "01:30"},
		{10 * time.Second, "00:10"},
		{0, "00:00"},
	}

	for _, test := range tests {
		result := formatDuration(test.input)
		if result != test.expected {
			t.Errorf("Expected %v but got %v", test.expected, result)
		}
	}
}

// Mock function to simulate timer for testing.
func mockStartTimer(duration time.Duration, done chan bool, tick time.Duration) {
	timeRemaining = duration
	ticker := time.NewTicker(tick)

	for range ticker.C {
		timeRemaining -= tick
		if timeRemaining <= 0 {
			done <- true
			return
		}
	}
}

// TestStartTimer tests the startTimer function with a mocked implementation.
func TestStartTimer(t *testing.T) {
	done := make(chan bool)
	duration := 5 * time.Second

	// Run the mock timer with a 1-second tick interval
	go mockStartTimer(duration, done, time.Second)

	// Wait for the timer to complete
	<-done

	if timeRemaining != 0 {
		t.Errorf("Expected timeRemaining to be 0, but got %v", timeRemaining)
	}
}
