package main

import (
	"fmt"
	"os"

	"github.com/samonzeweb/argononefan"
)

func main() {
	temperature, err := argononefan.ReadCPUTemperature()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("CPU temperature : %2.1f °C\n", temperature)
}
