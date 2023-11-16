package trayicon

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func New() *Ui {
	u := Ui{
		false,
		make(chan any),
	}
	go systray.Run(u.onReady, u.onExit)
	return &u
}

type Ui struct {
	active   bool
	quitChan chan any
}

func (u *Ui) OnQuit() <-chan any {
	return u.quitChan
}

func (u *Ui) ErrorMessage(fmt string, parts ...any) {
	dialog.Message(fmt, parts...).Error()
}

func (u *Ui) onReady() {
	systray.SetTitle("Logo Companion")
	systray.SetIcon(icon)
	settings := systray.AddMenuItem("Settings", "Open the settings menu")
	quit := systray.AddMenuItem("Quit", "Quit the companion")

	for {
		select {
		case <-settings.ClickedCh:
			openBrowser("http://localhost:3080")

		case <-quit.ClickedCh:
			u.quitChan <- "quit"
		}
	}
}

func (u *Ui) onExit() {}

func (u *Ui) Quit() {
	systray.Quit()
	time.Sleep(100 * time.Millisecond)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
