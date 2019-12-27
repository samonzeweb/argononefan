package argononefan

import (
	"fmt"

	"gobot.io/x/gobot/platforms/raspi"
)

// Fan address on i2c bus
const fanAddress = 0x1A

func SetFanSpeed(speed int) error {

	if speed < 0 || speed > 100 {
		return fmt.Errorf("desired fan speed is out of range : %d", speed)
	}

	adapter := raspi.NewAdaptor()
	defer adapter.Finalize()

	defaulti2c := adapter.GetDefaultBus()
	conn, err := adapter.GetConnection(fanAddress, defaulti2c)
	if err != nil {
		return fmt.Errorf("can't connect to i2c bus : %w", err)
	}
	defer conn.Close()

	err = conn.WriteByte(byte(speed))
	if err != nil {
		return fmt.Errorf("can't write fan seed : %W", err)
	}

	return nil
}
