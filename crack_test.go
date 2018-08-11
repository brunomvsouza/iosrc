// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package iosrc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.bmvs.io/iosrc"
)

func TestCrack(t *testing.T) {
	var table = []struct {
		key  string
		salt string
		out  string
		eout error
	}{
		{"aI9gkP3NEG+17D8UPyAT1ehRd7g=", "MTIzNDU", "", iosrc.ErrInvalidSalt},
		{"MTIzNDU=", "MTIzNDU=", "", iosrc.ErrSorryDidntWork},
		{"MTIzNDU", "MTIzNDU=", "", iosrc.ErrSorryDidntWork},
		{"hpO5ww==", "aI9gkP3NEG+17D8UPyAT1ehRd7g=", "", iosrc.ErrSorryDidntWork},
		{"aI9gkP3NEG+17D8UPyAT1ehRd7g=", "hpO5ww==", "4212", nil},
	}

	for _, row := range table {
		_, err := iosrc.Crack(row.key, row.salt)
		assert.Equal(t, err, row.eout)
	}
}
