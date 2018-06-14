package respwd

import (
	"crypto/sha1"
	"encoding/base64"
	"sync"

	"golang.org/x/crypto/pbkdf2"
)

const (
	keyIter = 1000
	keyLen  = 20
)

var (
	hash = sha1.New
)

func matchAsync(key string, salt, password []byte, wg *sync.WaitGroup,
	out chan []byte) {
	defer wg.Done()

	if match(key, salt, password) {
		out <- password
	}
}

func match(key string, salt, password []byte) bool {
	derivationKey := pbkdf2.Key(password, salt, keyIter, keyLen, hash)
	desiredKey := base64.StdEncoding.EncodeToString(derivationKey)
	return key == desiredKey
}
