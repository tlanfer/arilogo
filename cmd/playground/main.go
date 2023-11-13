package main

import (
	"api/internal/adapter/outbound/wled"
)

func main() {

	//instances := make(chan string)
	//go discovery.DiscoverWled(5*time.Second, instances)
	//
	//for instance := range instances {
	//	log.Println("Found", instance)

	client := wled.New("192.168.1.173")
	client.Effects()

	//}
}
