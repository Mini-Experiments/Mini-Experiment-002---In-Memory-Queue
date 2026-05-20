package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Event struct {
	ProducerID    string
	ConsumerID    string
	IPAddress     string
	PushTimestamp time.Time
	EventID       string
	Payload       string
}

type Node struct {
	Value Event
	Next  *Node
}

var head *Node
var mu sync.Mutex

func consumer1(event Event) {
	fmt.Println("This is consumer one.")
	fmt.Printf("The event details received are -- \nProducerID: %s\nConsumerID: %s\nIPAddress: %s\nTimestamp: %v\nEventID: %s\nPayload: %s\n\n",
		event.ProducerID, event.ConsumerID, event.IPAddress, event.PushTimestamp, event.EventID, event.Payload)
}

func consumer2(event Event) {
	fmt.Println("This is consumer two.")
	fmt.Printf("The event details received are -- \nProducerID: %s\nConsumerID: %s\nIPAddress: %s\nTimestamp: %v\nEventID: %s\nPayload: %s\n\n",
		event.ProducerID, event.ConsumerID, event.IPAddress, event.PushTimestamp, event.EventID, event.Payload)
}

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

func pop() {
	mu.Lock()
	if head == nil{
		mu.Unlock()
		return
	}
	
	temp := head
	head = head.Next
	mu.Unlock()

	// check for consumer ID
	if temp.Value.ConsumerID == "X" {
		go consumer1(temp.Value)
	} else {
		go consumer2(temp.Value)
	}
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

func producer(producerID string, wg *sync.WaitGroup) {
	defer wg.Done()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand1 := r.Intn(10) + 1
	localVariable := 0

	for i := 0; i < rand1; i++ {
		eventID := fmt.Sprintf("%s%d", producerID, localVariable)
		localVariable++

		event := Event{
			EventID: eventID,
		}

		random(&event)         
		event.ProducerID = producerID

		mu.Lock()
		push(event)
		mu.Unlock()

		time.Sleep(time.Duration(r.Intn(3000)+1000) * time.Millisecond)
	}
}

func main() {

	var wg sync.WaitGroup

	producerID := 'A'
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go producer(string(producerID), &wg)
		producerID++
	}

	// The pop function must continously check the queue.
	go func() {
		for {
			pop()
			time.Sleep(500 * time.Millisecond) // TO provide for some delay
		}
	}()

	wg.Wait()
}