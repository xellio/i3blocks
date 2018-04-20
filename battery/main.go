package main

import (
	"fmt"
	"log"
	"os/exec"
)

var (
	iconFull    = ""
	iconGood    = ""
	iconOkay    = ""
	iconPoor    = ""
	iconLow     = ""
	iconVeryLow = ""
	cFull       = "#FF0000"
	cGood       = "#FF0000"
	cOkay       = "#FF0000"
	cPoor       = "#FF6600"
	cLow        = "#FF0000"
	cVeryLow    = "#FF0000"
)

func main() {
	out, err := exec.Command("acpi", "b").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
