package encoding

import (
	"crypto/sha256"
	"testing"
)

func TestI2PDecode(t *testing.T) {
	s := "OUIPje~xpi7p~oadvhdo2OzCxLhVasobe0KYPic3fo8iWpnvXNj2ikgFz5EmAdD0gpWaQMYC3ghX0R3RiJXr~0ZV~qanunywMh9uANcuZp3GcA0kk6-GhfrMDImcGGBoiLUbcV6n5066tMV4RpJfPam2taLGm3HyutEayqqpRsk0XL98BQGJukH-xNLK2jxx5iv2UZV3IUtNpZEO8uzPnUyD6vYJDh32oiZCTWCUTDh~TTkGwimHVTo1hodhZ75s4MIanMC~QOQUJvw3Nq6nTIp-Jn6dgzWgxtf6XMO03D~OezK0svAQTO9QgqTmE1UXmqbzGX9Y0NWAINcdpv2vkYHVJ5F5n91~AVWR~U3JMwuxULo7TRSvJjpzUGvmEQkE5GZHblj3JV4JuA6BZDxK01KkdqgJ9V4gbtLyRnfEy0SIou-YsHsGbGIw6dRkhm8rzTNtfyo3FxllcutR3syn5onDUhVCUl837wHrYqJJJSupQRX29Q~BBnzbnlMfsaXVBQAEAAEAAA=="
	if len(s) <= 516 {
		t.Errorf("Incorrect input data length: %d less than 516", len(s))
	}
	data, err := Base64I2P.DecodeString(s)
	if err != nil {
		t.Errorf("%v is not nil", err)
	}

	// testing hashes
	h := sha256.Sum256(data)

	// b32
	hashed := Base32I2P.EncodeToString(h[:])
	expected := "cs3tmut6igkukwwyv6de42fblo7h4jecoav5gq2fzkfln3pr6ela"
	if hashed != expected {
		t.Errorf("Failed to test base 32; %s != %s", hashed, expected)
	}

	// b64
	hashed = Base64I2P.EncodeToString(h[:])
	expected = "FLc2Un5BlUVa2K-GTmihW75-JIJwK9NDRcqKtu3x8RY="
	if hashed != expected {
		t.Errorf("Failed to test base 64; %s != %s", hashed, expected)
	}
}
