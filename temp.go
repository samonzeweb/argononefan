package argononefan

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// File path in sysfs containing current CPU temperature
const temperatureFilePath = "/sys/class/thermal/thermal_zone0/temp"

// The temperature multiplier
const multiplier = float32(1000)

// ReadCPUTemperature reads the current CPU temperature
func ReadCPUTemperature() (float32, error) {
	content, err := ioutil.ReadFile(temperatureFilePath)
	if err != nil {
		return 0, fmt.Errorf("unable to read temperature file %s : %w", temperatureFilePath, err)
	}

	stringTemperature := strings.TrimSuffix(string(content), "\n")
	rawTemperature, err := strconv.Atoi(stringTemperature)
	if err != nil {
		return 0, fmt.Errorf("unable to parse temperature %s : %w", content, err)
	}

	return (float32(rawTemperature) / multiplier), nil
}
