package respwd_test

import (
	"testing"

	"github.com/brunomvsouza/iOSRestrictionCraker/pkg/respwd"
	"github.com/stretchr/testify/assert"
)

func TestCrack(t *testing.T) {
	var table = []struct {
		key  string
		salt string
		out  string
		eout error
	}{
		{"aI9gkP3NEG+17D8UPyAT1ehRd7g=", "MTIzNDU", "", respwd.ErrInvalidSalt},
		{"MTIzNDU=", "MTIzNDU=", "", respwd.ErrSorryDidntWork},
		{"MTIzNDU", "MTIzNDU=", "", respwd.ErrSorryDidntWork},
		{"hpO5ww==", "aI9gkP3NEG+17D8UPyAT1ehRd7g=", "", respwd.ErrSorryDidntWork},
		{"aI9gkP3NEG+17D8UPyAT1ehRd7g=", "hpO5ww==", "4212", nil},
	}

	for _, row := range table {
		_, err := respwd.Crack(row.key, row.salt)
		assert.Equal(t, err, row.eout)
	}
}
