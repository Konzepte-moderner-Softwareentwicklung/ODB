package odbclient

import "testing"

func TestCommand(t *testing.T) {
	command := OBDCommand{
		Mode: 0x01,
		PID:  0x01,
	}
	if command.String() != "0101" {
		t.Errorf("Expected '0101', got '%s'", command.String())
	}
}
