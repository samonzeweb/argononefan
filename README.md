# Argon One fan speed tools

## Project purpose

This project allows automatic fan speed adjustment of the Argon One Case
without Python. It is written in Go to use less memory (and for fun).

The project creates 3 tools :

* `readtemp` : display the current CPU temperature
* `setfan` : set the fan speed
* `adjustfan` : automatic fan speed adjustment

Both `setfan` and `adjustfan` require root privileges to access I2C device.

## Tested on...

I created this for my personnal use, and only tested on a Raspberry Pi 4 using Ubuntu 19.10 (arm64).

As the code use [gobot](https://gobot.io/) to access I2C device, it should run on other compatible OS.

The deploy part is rather *quick & dirty*. But it works.

## Build and install

To build, and install the tools (including starting service) :

```
make
sudo make install
```

The tools are installed in `/opt/argononefan`. After that `adjustfan` should run as a service. You can change thresholds (see below).

## Uninstall

```
sudo make uninstall
```

## Change fan thresholds

The file `/opt/argononefan/adjustfan.json` contains the configuration used to adjust the fan speed according to the CPU temperature.

You can change its content, and restart the service with :

```
sudo systemctl restart adjustfan.service
```

An example :

```json
{
    "thresholds": [
        {
            "temperature": 65,
            "fanspeed": 100
        },
        {
            "temperature": 60,
            "fanspeed": 50
        },
        {
            "temperature": 55,
            "fanspeed": 10
        }
    ]
}
```

Thresholds have to be ordered from the higher to lower temperature. Under the lowest temperature the fan is stopped.

## Licence

Released under the MIT License, see LICENSE.txt for more informations.==)