package main

import (
	"api/internal/adapter/inbound/api/effects"
	"api/internal/adapter/inbound/api/idle"
	patternById "api/internal/adapter/inbound/api/pattern/byId"
	patternList "api/internal/adapter/inbound/api/pattern/list"
	subById "api/internal/adapter/inbound/api/subs/byId"
	subList "api/internal/adapter/inbound/api/subs/list"
	"api/internal/adapter/inbound/api/test"
	"api/internal/adapter/outbound/fileConfig"
	"api/internal/adapter/outbound/wled"
	"api/internal/core/controller"
	"errors"
	"flag"
	"net/http"
)

var (
	filename = flag.String("c", "config.yaml", "filename for the config file")
)

func main() {
	flag.Parse()

	repo := fileConfig.New(*filename)

	wl := wled.New("192.168.1.173")

	c := controller.New(repo, wl)

	mux := http.NewServeMux()
	mux.Handle("/api/effects", effects.NewHandler(&wl))
	mux.Handle("/api/test", test.NewHandler(c))
	mux.Handle("/api/idle", idle.NewHandler(repo, c))

	mux.Handle("/api/pattern", patternList.NewHandler(repo))
	mux.Handle("/api/pattern/", http.StripPrefix("/api/pattern/", patternById.NewHandler(repo)))

	mux.Handle("/api/subs", subList.NewHandler(repo, &wl))
	mux.Handle("/api/subs/", http.StripPrefix("/api/subs/", subById.NewHandler(repo)))

	err := http.ListenAndServe(":3080", mux)
	if !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}
