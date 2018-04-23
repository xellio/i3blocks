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

	title := fmt.Sprintf("%s %d", time.Now().Month().String(), time.Now().Year())

	path, err := exec.LookPath("yad")
	if err != nil {
		path, err = exec.LookPath("notify-send")
		if err != nil {
			panic(err)
		}
		err = notify(path, title, cal)
		if err != nil {
			panic(err)
		}
	} else {
		err = notify(path, "--no-buttons", "--mouse", "--close-on-unfocus", "--skip-taskbar", "--calendar")
		if err != nil {
			panic(err)
		}
	}

}

func notify(path string, args ...string) error {
	cmd := exec.Command(path, args...)

	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
