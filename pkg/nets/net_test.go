package nets

import (
	"testing"
)

func TestGetIPv4ByInterface(T *testing.T) {
	ips, err := GetIPv4ByInterface("eth2")
	T.Errorf("ipv4s: %v,  error: %v", ips, err)
}

func TestLocalIPv4s(T *testing.T) {
	ips, err := LocalIPv4s()
	T.Errorf("ipv4s: %v,  error: %v", ips, err)
}
