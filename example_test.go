// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package iosrc_test

import (
	"fmt"

	"go.bmvs.io/iosrc"
)

func ExampleCrack() {
	key := "aI9gkP3NEG+17D8UPyAT1ehRd7g="
	salt := "hpO5ww=="
	pwd, err := iosrc.Crack(key, salt)
	if err != nil {
		panic(err)
	}

	fmt.Println(pwd)

	// Output: 4212
}
