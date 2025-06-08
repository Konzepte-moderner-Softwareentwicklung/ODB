package odbclient

import (
	"bufio"
	"errors"
	"net"
	"strings"
	"sync"
)

// ODBClient represents an OBD-II client
type ODBClient struct {
	conn   net.Conn
	reader *bufio.Reader
	mu     sync.Mutex
}

// NewODBClient creates a new OBD client connected to the specified address
func NewODBClient(addr string) (*ODBClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &ODBClient{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}, nil
}

// Close closes the connection
func (c *ODBClient) Close() error {
	return c.conn.Close()
}

func (c *ODBClient) Send(command OBDCommand) (string, error) {
	return c.SendCommand(command.String())
}

// SendCommand sends an OBD command and returns the response
func (c *ODBClient) SendCommand(command string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Send command
	cmd := command + "\r"
	if _, err := c.conn.Write([]byte(cmd)); err != nil {
		return "", err
	}

	// Read response
	response, err := c.reader.ReadString('>')
	if err != nil {
		return "", err
	}

	// Clean response
	response = strings.ReplaceAll(response, "\r", "")
	response = strings.ReplaceAll(response, "\n", "")
	response = strings.ReplaceAll(response, ">", "")
	response = strings.TrimSpace(response)

	// Handle error responses
	if strings.Contains(response, "?") {
		return "", errors.New("invalid command")
	}
	if strings.Contains(response, "NO DATA") {
		return "", errors.New("no data available")
	}

	return response, nil
}

// connect to OBD-II device
func (c *ODBClient) Connect() error {
	// Initialize connection
	_, err := c.SendCommand("ATZ")
	if err != nil {
		return err
	}

	_, err = c.SendCommand("ATE0")
	return err
}
