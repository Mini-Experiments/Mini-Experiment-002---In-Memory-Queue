package main

import (
	"fmt"
	"time"
)

type Event struct {
	ProducerID string
	ConsumerID string
	IPAddress  string
	Timestamp  time.Time
	EventID    string
	Payload    string
}

func pull(number int) Event {
	events := []Event{
		{ProducerID: "A", ConsumerID: "X", IPAddress: "192.168.00.1", Timestamp: time.Now().Add(4 * time.Hour), EventID: "A0", Payload: "payload-1"},
		{ProducerID: "A", ConsumerID: "Y", IPAddress: "192.168.00.2", Timestamp: time.Now().Add(3 * time.Hour), EventID: "A1", Payload: "payload-2"},
		{ProducerID: "B", ConsumerID: "X", IPAddress: "192.168.00.3", Timestamp: time.Now().Add(2 * time.Hour), EventID: "B0", Payload: "payload-3"},
		{ProducerID: "B", ConsumerID: "Y", IPAddress: "192.168.00.4", Timestamp: time.Now().Add(1 * time.Hour), EventID: "B1", Payload: "payload-4"},
	}
	return events[number-1]
}

func FirstConsumer(event Event) {
	fmt.Printf("This is the first consumer having received the event with the consumer ID %s, and event ID %s\n", event.ConsumerID, event.EventID)
}

func SecondConsumer(event Event) {
	fmt.Printf("This is the second consumer having received the event with the consumer ID %s, and event ID %s\n", event.ConsumerID, event.EventID)
}

func main() {
	number := 1
	for i := 1; i < 5; i++ {
		event := pull(number)

		if event.ConsumerID == "X" {
			FirstConsumer(event)
		} else {
			SecondConsumer(event)
		}
		number++
	}
}