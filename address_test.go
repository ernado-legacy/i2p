package i2p

import (
	"bytes"
	"net"
	"testing"
)

func TestDecodeDestination(t *testing.T) {
	s := "OUIPje~xpi7p~oadvhdo2OzCxLhVasobe0KYPic3fo8iWpnvXNj2ikgFz5EmAdD0gpWaQMYC3ghX0R3RiJXr~0ZV~qanunywMh9uANcuZp3GcA0kk6-GhfrMDImcGGBoiLUbcV6n5066tMV4RpJfPam2taLGm3HyutEayqqpRsk0XL98BQGJukH-xNLK2jxx5iv2UZV3IUtNpZEO8uzPnUyD6vYJDh32oiZCTWCUTDh~TTkGwimHVTo1hodhZ75s4MIanMC~QOQUJvw3Nq6nTIp-Jn6dgzWgxtf6XMO03D~OezK0svAQTO9QgqTmE1UXmqbzGX9Y0NWAINcdpv2vkYHVJ5F5n91~AVWR~U3JMwuxULo7TRSvJjpzUGvmEQkE5GZHblj3JV4JuA6BZDxK01KkdqgJ9V4gbtLyRnfEy0SIou-YsHsGbGIw6dRkhm8rzTNtfyo3FxllcutR3syn5onDUhVCUl837wHrYqJJJSupQRX29Q~BBnzbnlMfsaXVBQAEAAEAAA=="
	dest, err := DecodeDestination(s)
	if err != nil {
		t.Errorf("Failed to decode destination: %v", err)
	}

	if len(dest) < 300 {
		t.Errorf("Bad length of destination %d", len(dest))
	}
}

func TestBase32DestinationHash(t *testing.T) {
	s := "OUIPje~xpi7p~oadvhdo2OzCxLhVasobe0KYPic3fo8iWpnvXNj2ikgFz5EmAdD0gpWaQMYC3ghX0R3RiJXr~0ZV~qanunywMh9uANcuZp3GcA0kk6-GhfrMDImcGGBoiLUbcV6n5066tMV4RpJfPam2taLGm3HyutEayqqpRsk0XL98BQGJukH-xNLK2jxx5iv2UZV3IUtNpZEO8uzPnUyD6vYJDh32oiZCTWCUTDh~TTkGwimHVTo1hodhZ75s4MIanMC~QOQUJvw3Nq6nTIp-Jn6dgzWgxtf6XMO03D~OezK0svAQTO9QgqTmE1UXmqbzGX9Y0NWAINcdpv2vkYHVJ5F5n91~AVWR~U3JMwuxULo7TRSvJjpzUGvmEQkE5GZHblj3JV4JuA6BZDxK01KkdqgJ9V4gbtLyRnfEy0SIou-YsHsGbGIw6dRkhm8rzTNtfyo3FxllcutR3syn5onDUhVCUl837wHrYqJJJSupQRX29Q~BBnzbnlMfsaXVBQAEAAEAAA=="
	dest, err := DecodeDestination(s)
	if err != nil {
		t.Errorf("Failed to decode destination: %v", err)
	}
	expected := "cs3tmut6igkukwwyv6de42fblo7h4jecoav5gq2fzkfln3pr6ela"
	got := dest.Base32()
	if expected != got {
		t.Errorf("Bad b32 %s for destination; expected %s", got, expected)
	}
}

func TestIPv6Address(t *testing.T) {
	s := "OUIPje~xpi7p~oadvhdo2OzCxLhVasobe0KYPic3fo8iWpnvXNj2ikgFz5EmAdD0gpWaQMYC3ghX0R3RiJXr~0ZV~qanunywMh9uANcuZp3GcA0kk6-GhfrMDImcGGBoiLUbcV6n5066tMV4RpJfPam2taLGm3HyutEayqqpRsk0XL98BQGJukH-xNLK2jxx5iv2UZV3IUtNpZEO8uzPnUyD6vYJDh32oiZCTWCUTDh~TTkGwimHVTo1hodhZ75s4MIanMC~QOQUJvw3Nq6nTIp-Jn6dgzWgxtf6XMO03D~OezK0svAQTO9QgqTmE1UXmqbzGX9Y0NWAINcdpv2vkYHVJ5F5n91~AVWR~U3JMwuxULo7TRSvJjpzUGvmEQkE5GZHblj3JV4JuA6BZDxK01KkdqgJ9V4gbtLyRnfEy0SIou-YsHsGbGIw6dRkhm8rzTNtfyo3FxllcutR3syn5onDUhVCUl837wHrYqJJJSupQRX29Q~BBnzbnlMfsaXVBQAEAAEAAA=="
	dest, err := DecodeDestination(s)
	if err != nil {
		t.Errorf("Failed to decode destination: %v", err)
	}
	expected := net.ParseIP("9e5c:eafc:fd28:d9cb:6767:8179:7864:d81b")
	got := dest.Address()
	if !bytes.Equal(expected, got) {
		t.Errorf("Bad b32 %s for destination; expected %s", got, expected)
	}
}
