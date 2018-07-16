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
