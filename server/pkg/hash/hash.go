package hash

import (
	"crypto/sha1"
	"encoding/base64"
)

func Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
