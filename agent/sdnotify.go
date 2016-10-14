package main

// Code forked from github.com/coreos/go-systemd/daemon

import (
	"errors"
	"net"
	"os"
)

var sdNotifyNoSocket = errors.New("No socket")

// sdNotify sends a message to the init daemon. It is common to ignore the error.
func sdNotify(state string) error {
	socketAddr := &net.UnixAddr{
		Name: os.Getenv("NOTIFY_SOCKET"),
		Net:  "unixgram",
	}

	if socketAddr.Name == "" {
		return sdNotifyNoSocket
	}

	conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(state))
	return err
}
