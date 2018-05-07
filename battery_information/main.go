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

	output = appendBatteryInformation(output, information.BatteryInformation)
	output = appendAdapterInformation(output, information.AdapterInformation)
	output = appendThermalInformation(output, information.ThermalInformation)

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

func appendBatteryInformation(output string, information []*acpi.BatteryInformation) string {

	if len(information) > 1 {
		for _, bi := range information {
			if bi.Status != "Unknown" {
				output = fmt.Sprintf("%s\nBattery %d: %s\nLevel: %d", output, bi.Number, bi.Status, bi.Level)
			} else {
				output = fmt.Sprintf("%s\nBattery %d\nLevel: %d", output, bi.Number, bi.Level)
			}
		}
		return output + "\n"
	}

	return fmt.Sprintf("%s\nLevel: %d\n", output, information[0].Level)
}

func appendAdapterInformation(output string, information []*acpi.AdapterInformation) string {
	if len(information) > 1 {
		for _, ai := range information {
			output = fmt.Sprintf("%s\nAdapter %d: %s", output, ai.Number, ai.Status)
		}
		return output + "\n"
	}
	return fmt.Sprintf("%s\nAdapter: %s\n", output, information[0].Status)
}

func appendThermalInformation(output string, information []*acpi.ThermalInformation) string {
	if len(information) > 1 {
		for _, ti := range information {
			output = fmt.Sprintf("%s\nThermal %d: %.2f%s", output, ti.Number, ti.Degree, ti.Unit)
		}
		return output + "\n"
	}
	return fmt.Sprintf("%s\nThermal: %.2f%s", output, information[0].Degree, information[0].Unit)
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
