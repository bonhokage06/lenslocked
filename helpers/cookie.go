package helpers

import "encoding/base64"

func Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func Decode(value string) string {
	decoded, _ := base64.StdEncoding.DecodeString(value)
	return string(decoded)
}
