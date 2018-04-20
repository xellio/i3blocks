package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
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
	out, err := exec.Command("acpi", "b").Output()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	lines := bytes.Split(out[:len(out)-1], []byte("\n"))

	var fullText string
	var shortText string
	var color string

	for _, bat := range lines {
		values := bytes.Split(bat, []byte(":"))
		batPercent := strings.Trim(string(values[1][len(values[1])-4:]), " %")
		percent, err := strconv.Atoi(batPercent)
		if err != nil {
			continue
		}
		icon, color := iconAndColor(percent)

		fullText = fmt.Sprintf("%s <span foreground=\"%s\">%s %s%s</span>", fullText, color, icon, batPercent, "%")
	}

	fmt.Println(fullText)
	fmt.Println(shortText)
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
