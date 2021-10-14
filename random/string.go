package random

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

func GenerateSafeString(length int) (string, error) {
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		ret[i] = letters[RandomInt32Safe(0, len(letters))]
	}

	return string(ret), nil
}

func GenerateString(length int) string {
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		ret[i] = letters[RandomInt(0, len(letters))]
	}

	return string(ret)
}
