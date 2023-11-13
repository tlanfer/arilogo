package discovery

import (
	"api/internal/adapter/outbound/wled"
	"context"
	"encoding/json"
	"fmt"
	"github.com/grandcat/zeroconf"
	"log"
	"net/http"
	"time"
)

func DiscoverWled(timeout time.Duration, instances chan<- string) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			host := entry.AddrIPv4[0].String()
			go func() {
				if isValidDevice(host) {
					instances <- host
				}
			}()
		}
		log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err = resolver.Browse(ctx, "_http._tcp", "local.", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()

	close(instances)
}

func isValidDevice(host string) bool {
	jsonUrl := fmt.Sprintf("http://%v/json/info", host)
	//log.Printf("Check jsonUrl %v", jsonUrl)
	resp, err := http.Get(jsonUrl)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	dto := wled.Info{}
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return false
	}

	if dto.Brand == "WLED" {
		return true
	}

	return false
}
