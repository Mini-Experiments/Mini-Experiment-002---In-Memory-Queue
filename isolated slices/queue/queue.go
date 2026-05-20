package main

import (
	"fmt"
	"time"
)

// Event structure defined
type Event struct {
	ProducerID string
	ConsumerID string
	IPAddress  string
	Timestamp  time.Time
	EventID    string
	Payload    string
}

// Node structure defined
type Node struct {
	Value Event
	Next  *Node
}

// Global variable
var head *Node

// Push function
func push(event Event) {
	newNode := &Node{
		Value: event,
		Next:  nil,
	}

	if head == nil {
		head = newNode
	} else {
		temp := head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newNode
	}
}

// Pop function
func pop() {
	if head == nil {
		fmt.Println("The queue is already empty.")
	} else {
		temp := head
		head = head.Next
		fmt.Printf("Event has been popped: %s\n", temp.Value.EventID)
	}
}

func main() {
	// Hardcoded 4 events
	events := []Event{
		{ProducerID: "A", ConsumerID: "X", IPAddress: "192.168.1.1", Timestamp: time.Now().Add(4 * time.Hour), EventID: "A0", Payload: "payload-1"},
		{ProducerID: "A", ConsumerID: "Y", IPAddress: "192.168.1.2", Timestamp: time.Now().Add(3 * time.Hour), EventID: "A1", Payload: "payload-2"},
		{ProducerID: "B", ConsumerID: "X", IPAddress: "192.168.1.3", Timestamp: time.Now().Add(2 * time.Hour), EventID: "B0", Payload: "payload-3"},
		{ProducerID: "B", ConsumerID: "Y", IPAddress: "192.168.1.4", Timestamp: time.Now().Add(1 * time.Hour), EventID: "B1", Payload: "payload-4"},
	}

	var input int

	for {
		fmt.Print("Enter number (0-3 to push, -2 to pop, -1 to exit): ")
		fmt.Scanln(&input)

		if input == -1 {
			break
		} else if input == -2 {
			pop()
		} else if input >= 0 && input <= 3 {
			push(events[input])
			fmt.Printf("Event %s has been pushed onto the queue\n", events[input].EventID)
		} else {
			fmt.Println("Invalid input. Try again.")
		}
	}
}