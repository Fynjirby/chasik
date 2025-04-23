package main

import (
	"time"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		red := color.New(color.FgRed).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		magenta := color.New(color.FgHiMagenta).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		currectTime := time.Now()

		formatedH := currectTime.Format("15")
		formatedM := currectTime.Format("04")
		formatedS := currectTime.Format("05")
		formattedD := currectTime.Format("2")
		formattedMo := currectTime.Format("Jan")
		formattedY := currectTime.Format("2006")

		fmt.Printf("\r\t %s %s %s %s:%s:%s", cyan(formattedY), yellow(formattedMo), magenta(formattedD), red(formatedH), blue(formatedM), green(formatedS))
		<-ticker.C
	}
}
