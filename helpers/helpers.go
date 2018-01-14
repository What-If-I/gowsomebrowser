package helpers

import (
	"crypto/rand"
	"fmt"
)

// Simple pseudo uuid4 generator. Don't want to include any external dependencies
// Thanks to https://stackoverflow.com/a/25736155/7645774
func GenUUID4() (uuid string) {
	b := make([]byte, 16)
	rand.Read(b)
	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
