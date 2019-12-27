package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/samonzeweb/argononefan"
)

func main() {
	if len(os.Args) != 2 {
		displayUsageAndExit()
	}

	fanspeed, err := strconv.Atoi(os.Args[1])
	if err != nil || fanspeed < 0 || fanspeed > 100 {
		displayUsageAndExit()
	}

	err = argononefan.SetFanSpeed(fanspeed)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func displayUsageAndExit() {
	fmt.Fprintf(os.Stderr, "usage : %s fanspeed\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "with fanspeed between 0 and 100")
	os.Exit(1)
}
