package respwd

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"sync"

	"github.com/pkg/errors"
)

var (
	// ErrInvalidSalt couldn't parse salt
	ErrInvalidSalt = errors.New("respwd: couldn't parse salt")
	// ErrSorryDidntWork couldn't crack password
	ErrSorryDidntWork = errors.New("respwd: couldn't crack password")
)

// Crack brute forces iOS restriction password
func Crack(key, salt string) (string, error) {
	decodedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", ErrInvalidSalt
	}

	var wg sync.WaitGroup
	out := make(chan []byte, 1)
	for i := range rand.Perm(10000) {
		wg.Add(1)
		go matchAsync(key, decodedSalt, []byte(fmt.Sprintf("%04d", i)), &wg, out)
	}

	go func(wg *sync.WaitGroup, out chan []byte) {
		wg.Wait()
		close(out)
	}(&wg, out)

	password := <-out
	if len(password) == 0 {
		return "", ErrSorryDidntWork
	}

	return string(password), nil
}
