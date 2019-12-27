package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/samonzeweb/argononefan"
)

// Scan temperature (and adust fan speed) with the given internval
const adjustInterval = 5 * time.Second

// The fan speed is maintained for at least X intervals
// ie if interval is 5 seconds, and interval count is equal to 3, then
// the fan will not slow down for at least 15 secondes (5 * 3).
// This will not prevent the fan to speed up.
const maintainSpeedInIntervalCount = 12

// Configuration file (in current directory)
const configurationFile = "adjustfan.json"

func main() {
	configuration, err := readConfiguration(configurationFile)
	if err != nil {
		dislayErrorAndExit(err)
	}

	var stopsig = make(chan os.Signal, 1)
	signal.Notify(stopsig, syscall.SIGTERM)

	adjustFanLoop(configuration, stopsig)
	argononefan.SetFanSpeed(0)
}

func adjustFanLoop(configuration *Configuration, stopsig <-chan os.Signal) {
	previousFanSpeed := -1
	intervalsWithCurrentSpeed := 0
	for {
		cpuTemparature, err := argononefan.ReadCPUTemperature()
		if err != nil {
			dislayErrorAndExit(err)
		}

		fanSpeed := 0
		for _, t := range configuration.Thresholds {
			if cpuTemparature >= t.Temperature {
				fanSpeed = t.FanSpeed
				break
			}
		}

		if previousFanSpeed > 0 {
			intervalsWithCurrentSpeed++
		}

		if fanSpeed != previousFanSpeed {
			if fanSpeed > previousFanSpeed || (intervalsWithCurrentSpeed >= maintainSpeedInIntervalCount) {
				err := argononefan.SetFanSpeed(fanSpeed)
				if err != nil {
					dislayErrorAndExit(err)
				}
				previousFanSpeed = fanSpeed
				intervalsWithCurrentSpeed = 0
			}
		}

		select {
		case <-stopsig:
			return
		case <-time.After(adjustInterval):
			// nothing
		}

	}
}

func dislayErrorAndExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	_ = argononefan.SetFanSpeed(0)
	os.Exit(1)
}
