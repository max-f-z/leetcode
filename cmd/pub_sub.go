package main

import (
	"fmt"
	"sync"
)

type dataEvent struct {
	data  interface{}
	topic string
}

type subscriber chan dataEvent

type subscriberSlice []subscriber

type eventBus struct {
	subs  map[string]subscriberSlice
	mutex sync.RWMutex
}

func (eb *eventBus) subscribe(topic string, sub subscriber) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	if prev, ok := eb.subs[topic]; ok {
		eb.subs[topic] = append(prev, sub)
	} else {
		eb.subs[topic] = subscriberSlice{sub}
	}
}

func (eb *eventBus) publish(topic string, data interface{}) {
	eb.mutex.RLock()
	defer eb.mutex.RUnlock()

	if chans, ok := eb.subs[topic]; ok {
		chans = append(subscriberSlice{}, chans...)
		de := dataEvent{
			topic: topic,
			data:  data,
		}
		for _, v := range chans {
			v <- de
		}
	}
}

func printDataEvent(ch string, data dataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.topic, data.data)
}
