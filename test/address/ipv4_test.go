package address

import (
	"github.com/srogerf/ip_updater/address"
	"testing"
)

func TestGetIPv4(t *testing.T) {
	testMe := address.GetIPv4()
	if len(testMe) <= 0 {
		t.Error("No ip retuirned")
	}
}
