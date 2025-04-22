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

		currectTime := time.Now()

		formatedH := currectTime.Format("15")
		formatedM := currectTime.Format("04")
		formatedS := currectTime.Format("05")

		fmt.Printf("\r\t%s:%s:%s", red(formatedH), blue(formatedM), green(formatedS))
		<-ticker.C
	}
}
