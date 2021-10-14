package encoding

import "encoding/base64"

// BASE 64 ENCODE & DECODE

func Base64EncodeString(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func Base64EncodeBytes(input []byte) (bytes []byte) {
	bytes = make([]byte, 0)
	base64.StdEncoding.Encode(input, bytes)
	return
}

func Base64DecodeString(input string) (string, error) {
	if x, y := base64.StdEncoding.DecodeString(input); y != nil {
		return "", y
	} else {
		return string(x), nil
	}
}

func Base64DecodeBytes(input []byte) (bytes []byte) {
	bytes = make([]byte, 0)
	base64.StdEncoding.Decode(input, bytes)
	return
}
