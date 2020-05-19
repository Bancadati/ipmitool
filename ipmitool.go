package ipmitool

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	ipmitoolCommand = "ipmitool"
)

// NewClient returns a new IPMI client
// port 0 will use the default ipmi port (623)
func NewClient(addr string, port uint16, username, password string) (*Client, error) {
	if port == 0 {
		port = 623
	}
	cl := &Client{
		addr:     addr,
		port:     port,
		user:     username,
		password: password,
	}

	cl.Power = NewPower(cl)

	return cl, nil
}

// Client represents an IPMI client
type Client struct {
	addr     string
	port     uint16
	user     string
	password string
	Power    *Power
}

// getBaseParam returns the command parameters for the ipmitool command
// sets up the ipmitool command with host, user and password parameters
func (cl *Client) getBaseParam() []string {
	params := []string{
		"-H",
		cl.addr,
	}

	if cl.user != "" {
		params = append(params, "-U", cl.user)
	}
	if cl.password != "" {
		params = append(params, "-P", cl.password)
	}

	return params
}

// execute executes the provided command and returns the stdout, stderr and a potential error
// error is nil when command is successfully executed but returns an error condition.
// Revert to stderr for error response from the command
func (cl *Client) execute(args []string) (string, error) {
	cmd := exec.Command(ipmitoolCommand, args...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Failed to execute command: %w : %s", err, errBuf.String())
	}

	return outBuf.String(), nil
}
