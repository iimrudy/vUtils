package econding

import (
	"crypto/sha256"
	"crypto/sha512"
)

// SHA256 & SHA512 ENCODING

func Sha256EncodeBytes(input []byte) []byte {
	hash := sha256.Sum256(input)
	return hash[:]
}

func Sha256EncodeString(input string) string {
	return string(Sha256EncodeBytes([]byte(input)))
}

func Sha512EncodeBytes(input []byte) []byte {
	hash := sha512.Sum512(input)
	return hash[:]
}

func Sha512EncodeString(input string) string {
	return string(Sha512EncodeBytes([]byte(input)))
}
