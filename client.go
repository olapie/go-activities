package activities

import "net"

type Client struct {
	Application string
	DeviceID    string

	Addr net.Addr
}
