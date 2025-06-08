# ODB
Client for connecting to OBD devices like in cars used for communication with car's onboard diagnostic system.

## Installation of ODB emulator
```
  python3 -m pip install git+https://github.com/ircama/ELM327-emulator.git
```
## Usage of ODB emulator
Emulate a Toyota Auris Hybrid
```
  python3 -m elm -s car -n 35000
  CMD> scenario car
```
read more at https://github.com/ircama/ELM327-emulator
