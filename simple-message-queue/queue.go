package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Payload string
}

func (m *Message) execute() {
	m.Payload = fmt.Sprint(m.Payload, " is done")
}

type MQueue struct {
	queue    chan Message
	queueLen int
	wg       sync.WaitGroup
}

func newMqueue(qLen int) *MQueue {
	mq := &MQueue{
		queue: make(chan Message, qLen),
		queueLen: qLen,
	}

	go func() {
		for {
			task := mq.dequeue()
			task.execute()
			fmt.Println(task)
			time.Sleep(time.Second)
			mq.wg.Done()
		}
	}()

	return mq
}

func (mq *MQueue) enqueue(message Message) {
	fmt.Println(message.Payload, "is enqueued")
	mq.queue <- message
	mq.wg.Add(1)
}

func (mq *MQueue) dequeue() Message {
	return <-mq.queue
}

func main() {
	mqueue := newMqueue(5)
	m1 := Message{"task 1"}
	m2 := Message{"task 2"}
	m3 := Message{"task 3"}
	mqueue.enqueue(m1)
	mqueue.enqueue(m2)
	mqueue.enqueue(m3)
	mqueue.wg.Wait()
}
