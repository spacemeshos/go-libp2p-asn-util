package asnutil

import (
	"net"
)

var Store *disabled

// disabled builds the underlying trie on first call to AsnForIPv6.
// Alternatively, Init can be called to manually trigger initialization.
type disabled struct {
}

// AsnForIPv6 returns the AS number for the given IPv6 address.
// If no mapping exists for the given IP, this function will
// return an empty ASN and a nil error.
func (a *disabled) AsnForIPv6(ip net.IP) (string, error) {
	return "", nil
}
