package main

import (
	"errors"
	"net"
)

func externalIP() (string, error) {

	ifaces, err := net.Interfaces()

	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				if !v.IP.IsLoopback() {
					return v.IP.String(), nil
				}
			}
		}
	}

	return "", errors.New("are you connected to the network?")
}
