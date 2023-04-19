package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/bonhokage06/lenslocked/constants"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}
	if nRead != n {
		return nil, fmt.Errorf("could not read enough bytes")
	}
	return b, nil
}

// String returns a random string of length n
// n is the number of bytes to generate
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
func SessionToken() (string, error) {
	return String(constants.SessionTokenBytes)
}
