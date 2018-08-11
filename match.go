// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package iosrc

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
