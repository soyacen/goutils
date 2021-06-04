package netutils

import (
	"fmt"
	"net"
)

var (
	privateBlocks map[*net.IPNet]struct{}
)

func init() {
	privateBlocks = make(map[*net.IPNet]struct{})
	AppendPrivateBlocks(
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"100.64.0.0/10",
		"fd00::/8",
	)
}

// AppendPrivateBlocks append private network blocks
func AppendPrivateBlocks(bs ...string) {
	for _, b := range bs {
		if _, block, err := net.ParseCIDR(b); err == nil {
			privateBlocks[block] = struct{}{}
		}
	}
}

// IPs get all IP addresses
func IPs() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	var ips []string
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ip := AddrToIP(addr); ip != nil {
				ips = append(ips, ip.String())
			}
		}
	}
	return ips
}

// LocalIPs get all IP addresses of the local machine
func LocalIPs() []string {
	return IPs()
}

// AddrToIP Convert net.Addr to IP address
func AddrToIP(addr net.Addr) net.IP {
	switch v := addr.(type) {
	case *net.IPAddr:
		return v.IP
	case *net.IPNet:
		return v.IP
	case *net.TCPAddr:
		return v.IP
	case *net.UDPAddr:
		return v.IP
	default:
		return nil
	}
}

// IsPrivateIP Determine if the IP address is private
func IsPrivateIP(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return false
	}
	for block := range privateBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// IsLoopback reports whether ipAddr is a loopback address.
func IsLoopback(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return false
	}
	return ip.IsLoopback()
}

// IsUnspecified reports whether ipAddr is an unspecified address, either
// the IPv4 address "0.0.0.0" or the IPv6 address "::".
func IsUnspecified(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return false
	}
	return ip.IsUnspecified()
}

// RealIP get a public IP address
func RealIP(host string) (string, error) {
	// if addr specified then its returned
	if len(host) > 0 {
		if host != "0.0.0.0" && host != "[::]" && host != "::" {
			return host, nil
		}
	}
	var privateIPs []string
	var publicIPs []string
	var loopbackIPs []string
	for _, ipAddr := range IPs() {
		ip := net.ParseIP(ipAddr)
		if ip == nil {
			continue
		}
		if ip.IsUnspecified() {
			continue
		}
		if ip.IsLoopback() {
			loopbackIPs = append(loopbackIPs, ipAddr)
		} else if IsPrivateIP(ipAddr) {
			privateIPs = append(privateIPs, ipAddr)
		} else {
			publicIPs = append(publicIPs, ipAddr)
		}
	}
	if len(privateIPs) > 0 {
		return privateIPs[0], nil
	} else if len(publicIPs) > 0 {
		return publicIPs[0], nil
	} else if len(loopbackIPs) > 0 {
		return loopbackIPs[0], nil
	}
	return "", fmt.Errorf("no real IP address found")
}
