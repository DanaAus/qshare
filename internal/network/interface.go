package network

import (
	"errors"
	"net"
	"strings"
)

// GetActiveIPv4Interface returns the primary non-virtual IPv4 address.
func GetActiveIPv4Interface() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		// Skip down and loopback interfaces
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// Skip common virtual interface names (Docker, WSL, VirtualBox, VMware)
		name := strings.ToLower(iface.Name)
		if strings.Contains(name, "veth") ||
			strings.Contains(name, "docker") ||
			strings.Contains(name, "wsl") ||
			strings.Contains(name, "virtual") ||
			strings.Contains(name, "vmware") {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Looking for an IPv4 address
			if ip != nil && ip.To4() != nil {
				return ip.String(), nil
			}
		}
	}

	return "", errors.New("no active non-virtual network interface found")
}

// GetDisplayIP returns a fake IP if demo mode is active.
func GetDisplayIP(realIP string, demo bool) string {
	if demo {
		return "192.168.100.100"
	}
	return realIP
}

// GetDisplayURL replaces the real IP in a URL with a fake one if demo mode is active.
// It preserves the real port used in the URL.
func GetDisplayURL(realURL string, demo bool) string {
	if !demo {
		return realURL
	}
	
	// Find the scheme end
	start := strings.Index(realURL, "//")
	if start == -1 {
		return realURL
	}
	start += 2
	
	// Find the end of host:port part
	end := strings.Index(realURL[start:], "/")
	hostPort := ""
	if end == -1 {
		hostPort = realURL[start:]
	} else {
		hostPort = realURL[start : start+end]
	}

	// Extract port from hostPort if present
	port := "8080"
	if strings.Contains(hostPort, ":") {
		_, p, err := net.SplitHostPort(hostPort)
		if err == nil {
			port = p
		}
	}
	
	fakedHostPort := "192.168.100.100:" + port
	
	if end == -1 {
		return realURL[:start] + fakedHostPort
	}
	return realURL[:start] + fakedHostPort + realURL[start+end:]
}
