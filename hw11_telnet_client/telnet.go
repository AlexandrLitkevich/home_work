package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type ClientConnection struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &ClientConnection{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *ClientConnection) Connect() error {
	connection, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}

	c.connection = connection

	return nil
}

func (c *ClientConnection) Send() error {
	scanner := bufio.NewScanner(c.in)
	for scanner.Scan() {
		_, err := c.connection.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			return fmt.Errorf("fail to send: %w", err)
		}
	}
	return nil
}

func (c *ClientConnection) Receive() error {
	_, err := io.Copy(c.out, c.connection)

	if err != nil {
		return fmt.Errorf("fail to receive: %w", err)
	}
	return nil
}

func (c *ClientConnection) Close() error {
	return c.connection.Close()
}
