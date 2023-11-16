package main

import (
	bitsHandler "api/internal/adapter/inbound/api/bits"
	"api/internal/adapter/inbound/api/device"
	"api/internal/adapter/inbound/api/donations"
	giftsHandler "api/internal/adapter/inbound/api/gifts"
	"api/internal/adapter/inbound/api/idle"
	"api/internal/adapter/inbound/api/presets"
	resubsHandler "api/internal/adapter/inbound/api/resubs"
	"api/internal/adapter/inbound/api/streamlabshandler"
	"api/internal/adapter/inbound/api/test"
	"api/internal/adapter/inbound/api/twitchhandler"
	"api/internal/adapter/inbound/api/ui"
	twitchchat "api/internal/adapter/inbound/twitch"
	"api/internal/adapter/outbound/fileConfig"
	"api/internal/adapter/outbound/wled"
	"api/internal/core"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/tlanfer/go-streamlabs"
	"log"
	"net/http"
)

var (
	filename = flag.String("c", "config.yaml", "filename for the config file")
)

func main() {
	flag.Parse()

	config := fileConfig.New(*filename)

	wl, err := initLight(config)
	if err != nil {
		log.Printf("failed to start")
		return
	}

	twitch, err := initTwitch(config)
	if err != nil {
		log.Printf("failed to start: %v", err)
		return
	}

	sl, err := initStreamlabs(config)
	if err != nil {
		log.Printf("failed to start: %v", err)
		return
	}

	controller := core.NewController(config, wl)
	mananger := core.NewEventManager(controller, config, twitch, sl)
	go controller.Run()
	go mananger.Run()
	serve(wl, twitch, sl, controller, config)
}

func initLight(config core.GlobalConfigRepo) (core.Light, error) {
	addr, err := config.GetDeviceAddr(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to read device addr from config: %v", err)
	}

	wl := wled.New(addr)
	return &wl, nil
}

func initTwitch(config core.GlobalConfigRepo) (twitchchat.Twitch, error) {

	channel, err := config.GetChannel(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to read channel from config: %v", err)
	}

	twitch := twitchchat.New()

	if channel != "" {
		if err := twitch.Connect(channel); err != nil {
			return nil, fmt.Errorf("failed to connect to twitch: %v", err)
		}
	}

	return twitch, nil
}

func initStreamlabs(config core.GlobalConfigRepo) (streamlabs.Streamlabs, error) {
	token, err := config.GetStreamlabsToken(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to read streamlabs token from config: %v", err)
	}

	sl := streamlabs.New()
	if token != "" {
		if err := sl.Connect(token); err != nil {
			return nil, fmt.Errorf("failed to connect to streamlabs: %v", err)
		}
	}
	return sl, nil

}

func serve(light core.Light, tc twitchchat.Twitch, sl streamlabs.Streamlabs, c *core.Controller, config fileConfig.Repo) {
	mux := http.NewServeMux()

	mux.Handle("/api/idle", idle.NewHandler(config, c))
	mux.Handle("/api/presets", presets.NewHandler(light))
	mux.Handle("/api/device/scan", device.NewScanner())
	mux.Handle("/api/device/addr", device.NewDeviceHandler(config, light))
	mux.Handle("/api/test", test.NewHandler(c))
	mux.Handle("/api/twitch", twitchhandler.NewHandler(config, tc))
	mux.Handle("/api/streamlabs", streamlabshandler.NewHandler(config, sl))

	mux.Handle("/api/bits", http.StripPrefix("/api/bits", bitsHandler.List(config)))
	mux.Handle("/api/bits/", http.StripPrefix("/api/bits/", bitsHandler.ById(config)))

	mux.Handle("/api/resubs", http.StripPrefix("/api/resubs", resubsHandler.List(config)))
	mux.Handle("/api/resubs/", http.StripPrefix("/api/resubs/", resubsHandler.ById(config)))

	mux.Handle("/api/gifts", http.StripPrefix("/api/gifts", giftsHandler.List(config)))
	mux.Handle("/api/gifts/", http.StripPrefix("/api/gifts/", giftsHandler.ById(config)))

	mux.Handle("/api/donations", http.StripPrefix("/api/donations", donations.List(config)))
	mux.Handle("/api/donations/", http.StripPrefix("/api/donations/", donations.ById(config)))

	mux.Handle("/", ui.NewHandler())

	err := http.ListenAndServe(":3080", mux)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed:", err)
	}
}
