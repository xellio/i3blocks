package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/xellio/gocal"
)

func main() {
	calendar := gocal.Cal{
		NoFormat:   true,
		HideHeader: true,
	}
	cal, err := calendar.Output()
	if err != nil {
		panic(err)
	}

	path, err := exec.LookPath("notify-send")
	if err != nil {
		panic(err)
	}
	//  cmd := exec.Command(path, time.Now().Month().String(), cal)
	title := fmt.Sprintf("%s %d", time.Now().Month().String(), time.Now().Year())
	cmd := exec.Command(path, title, fmt.Sprint(cal))

	if err = cmd.Start(); err != nil {
		panic(err)
	}
	if err = cmd.Wait(); err != nil {
		panic(err)
	}
}
