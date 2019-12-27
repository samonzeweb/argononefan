package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	Thresholds []Threshold `json:"thresholds"`
}

type Threshold struct {
	Temperature float32 `json:"temperature"`
	FanSpeed    int     `json:"fanspeed"`
}

func readConfiguration(filePath string) (*Configuration, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read configuration file %s : %w", filePath, err)
	}

	configuration := &Configuration{}
	err = json.Unmarshal(content, configuration)
	if err != nil {
		return nil, fmt.Errorf("unable to parse configuration file %s : %w", filePath, err)
	}

	err = checkConfiguration(configuration)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration content in %s : %w", filePath, err)
	}

	return configuration, nil
}

func checkConfiguration(configuration *Configuration) error {
	if len(configuration.Thresholds) == 0 {
		return fmt.Errorf("configuration require at least one threshold")
	}

	previousTemperature := float32(999)
	for _, threshold := range configuration.Thresholds {
		if threshold.FanSpeed < 0 || threshold.FanSpeed > 100 {
			return fmt.Errorf("invalid fanspeed [0-100]")
		}
		if threshold.Temperature < 40 || threshold.Temperature > 80 {
			return fmt.Errorf("invalid temperature [40-80]")
		}
		if threshold.Temperature > previousTemperature {
			return fmt.Errorf("temperatures have to be declarared from higher to lower")
		}
		previousTemperature = threshold.Temperature
	}

	return nil
}
