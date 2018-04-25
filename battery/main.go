package main

import (
	"bytes"
	"fmt"
	"strconv"
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

	out, err := acpi.Battery()
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(out[:len(out)-1], []byte("\n"))

	var fullText string
	var shortText string
	var color string
	var discharging bool

	for _, bat := range lines {

		values := bytes.Split(bat, []byte(","))

		if (string(values[0][len(values[0])-11:])) == "Discharging" {
			discharging = true
		}

		batPercent := strings.Trim(string(values[1]), " %")
		percent, err := strconv.Atoi(batPercent)
		if err != nil {
			continue
		}
		icon, color := iconAndColor(percent)

		fullText = fmt.Sprintf("%s <span foreground=\"%s\">%s %s%s</span>", fullText, color, icon, batPercent, "%")
		shortText = fmt.Sprintf("%s <span foreground=\"%s\">%s%s</span>", shortText, color, batPercent, "%")
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
