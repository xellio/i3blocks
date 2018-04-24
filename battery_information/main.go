package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("acpi")
	if err != nil {
		panic(err)
	}

	out, err := exec.Command(path, "-V").Output()
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(out[:len(out)-1], []byte("\n"))

	for _, line := range lines {
		fmt.Println(string(line))
	}

	path, err = exec.LookPath("yad")
	if err != nil {
		path, err = exec.LookPath("notify-send")
		if err != nil {
			panic(err)
		}
		err = notify(path, "", string(out))
		if err != nil {
			panic(err)
		}
	} else {
		err = notify(path, "--no-buttons", "--mouse", "--close-on-unfocus", "--skip-taskbar", "--text="+string(out))
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
	_ = cmd.Wait()
	return nil
}
