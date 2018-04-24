package main

import (
	"os/exec"
)

func main() {
	path, err := exec.LookPath("yad")
	if err != nil {
		panic(err)
	}

	err = notify(path,
		"--no-buttons",
		"--mouse",
		"--close-on-unfocus",
		"--skip-taskbar",
		"--calendar",
	)
	if err != nil {
		panic(err)
	}
}

func notify(path string, args ...string) error {
	cmd := exec.Command(path, args...)

	err := cmd.Start()
	if err != nil {
		return err
	}
	_ = cmd.Wait()
	return nil
}
