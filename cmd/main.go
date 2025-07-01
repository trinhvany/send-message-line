package main

import (
	"small_demo_go/publisher"
	"small_demo_go/subscriber"
	"time"
)

func main() {
	publisher.StartPublisher()
	subcriber.StartSubcriber()
	time.Sleep(5 * time.Second)
}