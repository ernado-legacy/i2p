package encoding

import (
	"encoding/base32"
	"encoding/base64"
)

const i2pB64alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~"

// Base64I2P is base64 encoding that is used for destination keys in i2p
var Base64I2P = base64.NewEncoding(i2pB64alphabet)

type tBase32I2PEncoding struct{}

const base32I2PEncodingDic = "abcdefghijklmnopqrstuvwxyz234567"

var base32Encoding = base32.NewEncoding(base32I2PEncodingDic)

func (b tBase32I2PEncoding) EncodeToString(src []byte) (s string) {
	s = base32Encoding.EncodeToString(src)
	for s[len(s)-1] == '=' {
		s = s[:len(s)-1]
	}
	return s
}

// Base32I2P is base32 encoding that is used for *.b32.i2p names
// It is sha256(destination key) in lowercase base32 with no padding (= characters)
var Base32I2P = tBase32I2PEncoding{}
