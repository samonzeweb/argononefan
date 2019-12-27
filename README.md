# Argon One fan speed tools

## Project purpose

This project allows automatic fan speed adjustment of the Argon One Case
without Python. It is written in Go to use less memory (and for fun).

The project creates 3 tools :

* `readtemp` : display the current CPU temperature
* `setfan` : set the fan speed
* `adjustfan` : automatic fan speed adjustment

Both `setfan` and `adjustfan` require root privileges to access i2c device.

## Build and install

To build, and install the tools (including starting service) :

```
make
sudo make install
```

The tools are installer in `/opt/argononefan`
## Uninstall

```
sudo make uninstall
```

## Licence

Released under the MIT License, see LICENSE.txt for more informations.==)