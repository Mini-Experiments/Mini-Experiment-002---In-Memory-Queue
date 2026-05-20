package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Event struct {
	ProducerID string
	ConsumerID string
	IPAddress  string
	PushTimestamp  time.Time
	EventID    string
	Payload    string
}

func random(event *Event) {
	event.IPAddress = fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256),rand.Intn(256),rand.Intn(256),rand.Intn(256))
	event.Payload = fmt.Sprintf("payload-%d", rand.Intn(10000))
	if rand.Intn(2) == 0 {
		event.ConsumerID = "X"
	} else {
		event.ConsumerID = "Y"
	}
}

func push(event *Event) {
	event.PushTimestamp = time.Now()
	fmt.Printf("The event %s has been successfully pushed on the queue\n", event.EventID)
}

func process(producerID string, wg *sync.WaitGroup) {
	defer wg.Done()   // important

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand1 := rand.Intn(10) + 1
	localVariable := 0
	// looping event generation a random number of time.
	for i := 0; i < rand1; i++ {
		eventID := fmt.Sprintf("%s%d", producerID, localVariable)
		localVariable++

		event := Event{
			EventID:eventID,
		}

		random(&event)
		event.ProducerID = producerID
		push(&event)
		
		time.Sleep(time.Duration(r.Intn(3000)+1000) * time.Millisecond)
		
	}
	fmt.Printf("Total number of events created by %s are %d. \n",producerID,localVariable)
}

func main() {
	var wg sync.WaitGroup

	producerID := 'A'
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go process(string(producerID), &wg)
		producerID++
	}

	wg.Wait()   // waits until all goroutines call Done()
}