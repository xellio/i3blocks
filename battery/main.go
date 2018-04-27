package main

import (
	"fmt"
	"strings"

	"github.com/xellio/tools/acpi"
)

var (
	iconCharging = ""
	iconFull     = "" // >90
	iconGood     = "" // >80
	iconOkay     = "" // >50
	iconPoor     = "" // >20
	iconLow      = "" // <20
	cFull        = "#05F900"
	cGood        = "#BFF929"
	cOkay        = "#F9F420"
	cPoor        = "#F9C11A"
	cLow         = "#FF6600"
)

func main() {

	batteryInformation, err := acpi.Battery()
	if err != nil {
		panic(err)
	}

	var fullText string
	var shortText string
	var color string
	var discharging bool

	for _, battery := range batteryInformation {
		if battery.Status == "Discharging" {
			discharging = true
		}

		icon, color := iconAndColor(battery.Level)

		fullText = fmt.Sprintf("%s <span foreground=\"%s\">%s %d%s</span>", fullText, color, icon, battery.Level, "%")
		shortText = fmt.Sprintf("%s<span foreground=\"%s\">%d%s</span>", shortText, color, battery.Level, "%")

	}

	if !discharging {
		fullText = fmt.Sprintf("%s %s", iconCharging, fullText)
		shortText = fmt.Sprintf("%s %s", iconCharging, shortText)
	}

	fmt.Println(strings.TrimSpace(fullText))
	fmt.Println(strings.TrimSpace(shortText))
	fmt.Println(color)
}

//
// iconAndColor returns the icon and the color for the given percent value
//
func iconAndColor(percent int) (string, string) {
	switch {
	case percent > 90:
		return iconFull, cFull
	case percent > 80:
		return iconGood, cGood
	case percent > 50:
		return iconOkay, cOkay
	case percent > 20:
		return iconPoor, cPoor
	default:
		return iconLow, cLow
	}
}
