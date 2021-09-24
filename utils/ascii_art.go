package utils

import (
	"fmt"

	figure "github.com/common-nighthawk/go-figure"
)

// PrintServer does ascii art for the string server
func PrintServer() {
	myFigure := figure.NewFigure("Zovino Server", "", true)
	myFigure.Print()
	fmt.Println()
}

// PrintClient does ascii art for string client
func PrintClient() {
	myFigure := figure.NewFigure("Zovino Client", "", true)
	myFigure.Print()
	fmt.Println()
}
