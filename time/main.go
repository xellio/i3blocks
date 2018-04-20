package main

import (
	"fmt"
	"time"
)

var (
	icon         = ""
	night        = "#A1A1A1"
	earlyMorning = "#42f44e"
	morning      = "#2a8230"
	forenoon     = "#11AAAA"
	noon         = "#808229"
	afternoon    = "#FF6600"
	evening      = "#FF0000"
	late         = "#11AAAA"
)

func main() {

	now := time.Now()
	hour := now.Hour()

	var color string

	switch hour {
	case 0, 1, 2, 3, 4:
		color = night
		break
	case 5, 6, 7, 8:
		color = morning
	case 9, 10, 11:
		color = forenoon
	case 12, 13, 14:
		color = noon
	case 15, 16:
		color = afternoon
	case 17, 18, 19:
		color = evening
		break
	case 20, 21, 22, 23:
		color = late
		break
	default:
		color = night
	}

	fmt.Println(icon + " " + now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(color)
}
