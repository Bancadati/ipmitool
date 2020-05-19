package ipmitool

import (
	"fmt"
	"strings"
)

// PowerState represents an IPMI node/server state of power
type PowerState string

const (
	// PowerStateOn represents a powered on state
	PowerStateOn PowerState = "on"
	// PowerStateOff represents a powered off state
	PowerStateOff PowerState = "off"
	// PowerStateUnknown represents an unknown power state
	PowerStateUnknown PowerState = "unknown"
)

// NewPower returns a new instance of the ipmi power sub command
func NewPower(cl *Client) *Power {
	return &Power{
		cl: cl,
	}
}

// Power represents the ipmi power command
type Power struct {
	cl *Client
}

// On sends a power on command to the ipmi server
func (p *Power) On() error {
	params := p.cl.getBaseParam()
	params = append(params, "power", "on")

	_, err := p.cl.execute(params)
	if err != nil {
		return fmt.Errorf("failed to power on: %w", err)
	}

	return nil
}

// Off sends a power off command to the ipmi server
func (p *Power) Off() error {
	params := p.cl.getBaseParam()
	params = append(params, "power", "off")

	_, err := p.cl.execute(params)
	if err != nil {
		return fmt.Errorf("failed to power on: %w", err)
	}

	return nil
}

// Cycle sends a power cycle command to the ipmi server
func (p *Power) Cycle() error {
	params := p.cl.getBaseParam()
	params = append(params, "power", "cycle")

	_, err := p.cl.execute(params)
	if err != nil {
		return fmt.Errorf("failed to power on: %w", err)
	}

	return nil
}

// Status fetches the current power state of the ipmi server
func (p *Power) Status() (PowerState, error) {
	params := p.cl.getBaseParam()
	params = append(params, "power", "status")

	stdout, err := p.cl.execute(params)
	if err != nil {
		return PowerStateUnknown, fmt.Errorf("failed to fetch power state: %w", err)
	}

	state := PowerStateUnknown
	if strings.Contains(stdout, "on") {
		state = PowerStateOn
	}
	if strings.Contains(stdout, "off") {
		state = PowerStateOff
	}

	return state, nil
}
