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

type threshold struct {
	temperature float32
	fanspeed    int
}

var thresholds = [...]threshold{
	threshold{temperature: 65, fanspeed: 100},
	threshold{temperature: 60, fanspeed: 50},
	threshold{temperature: 50, fanspeed: 10},
}

func main() {
	var stopsig = make(chan os.Signal, 1)
	signal.Notify(stopsig, syscall.SIGTERM)

	adjustFanLoop(stopsig)
	argononefan.SetFanSpeed(0)
}

func adjustFanLoop(stopsig <-chan os.Signal) {
	previousFanSpeed := -1
	intervalsWithCurrentSpeed := 0
	for {
		cpuTemparature, err := argononefan.ReadCPUTemperature()
		if err != nil {
			dislayErrorAndExit(err)
		}

		fanSpeed := 0
		for _, t := range thresholds {
			if cpuTemparature >= t.temperature {
				fanSpeed = t.fanspeed
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
