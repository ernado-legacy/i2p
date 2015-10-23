package i2p

import (
	"crypto/sha256"
	"io"
	"net"

	"github.com/cydev/i2p/encoding"
	"golang.org/x/crypto/sha3"
)

const destinationMinimumLength = 516

// DecodeDestination from i2p base64 destination key format
func DecodeDestination(s string) (d Destination, err error) {
	if len(s) < destinationMinimumLength {
		return nil, io.ErrUnexpectedEOF
	}
	data, err := encoding.Base64I2P.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return Destination(data), err
}

// Destination is public key for i2p router
type Destination []byte

// Address returns ipv6 ephemeral address based on sha3 hash
func (d Destination) Address() net.IP {
	var (
		h    = sha3.NewShake128()
		data = make([]byte, net.IPv6len)
	)
	h.Write(d)
	h.Read(data)
	return net.IP(data)
}

// Base32 is sha256 hash encoded in i2p b32
func (d Destination) Base32() string {
	data := sha256.Sum256(d)
	return encoding.Base32I2P.EncodeToString(data[:])
}