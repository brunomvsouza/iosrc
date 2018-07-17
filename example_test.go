package iosrc_test

import (
	"fmt"

	"go.bmvs.io/iosrc"
)

func ExampleCrack() {
	key := "aI9gkP3_invalid_T1ehRd7g="
	salt := "H_invalid_=="
	pwd, err := iosrc.Crack(key, salt)
	if err != nil {
		panic(err)
	}

	fmt.Println(pwd)

	// Output: 1234
}
