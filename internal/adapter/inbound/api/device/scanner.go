package device

import (
	"api/internal/adapter/outbound/wled/discovery"
	"fmt"
	"log"
	"net/http"
	"time"
)

func NewScanner() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}

		ch := make(chan string)
		go discovery.DiscoverWled(10*time.Second, ch)

		for row := range ch {
			log.Printf("Found WLED device at %v", row)
			_, _ = fmt.Fprintln(w, row)
			flusher.Flush()
		}
	})
}
