package main

import "testing"

func TestPubSub(t *testing.T) {
	ch1 := make(chan dataEvent)
	ch2 := make(chan dataEvent)
	ch3 := make(chan dataEvent)

	eb := &eventBus{
		subs: map[string]subscriberSlice{},
	}

	eb.subscribe("topic1", ch1)
	eb.subscribe("topic2", ch2)
	eb.subscribe("topic2", ch3)

	go eb.publish("topic1", "Hi topic 1")
	go eb.publish("topic2", "Welcome to topic 2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}
