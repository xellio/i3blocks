package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/xellio/tools/acpi"
)

var path string

func main() {

	path, err := exec.LookPath("yad")
	if err != nil {
		panic(err)
	}

	information, err := acpi.Everything()
	if err != nil {
		panic(err)
	}

	var output string

	for _, bi := range information.BatteryInformation {
		output = fmt.Sprintf("%s\nBattery %d: %s\nLevel: %d", output, bi.Number, bi.Status, bi.Level)
	}

	for _, ai := range information.AdapterInformation {
		output = fmt.Sprintf("%s\nAdapter %d: %s", output, ai.Number, ai.Status)
	}

	for _, ti := range information.ThermalInformation {
		output = fmt.Sprintf("%s\nThermal %d: %f%s", output, ti.Number, ti.Degree, ti.Unit)
	}

	err = notify(path,
		"--no-buttons",
		"--mouse",
		"--close-on-unfocus",
		"--skip-taskbar",
		"--text="+strings.TrimSpace(output),
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
