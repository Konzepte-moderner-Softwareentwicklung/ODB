package odbclient

import "fmt"

type OBDCommand struct {
	Mode uint8
	PID  uint8
}

func (cmd OBDCommand) String() string {
	return fmt.Sprintf("%02X", cmd.Mode) + fmt.Sprintf("%02X", cmd.PID)
}
