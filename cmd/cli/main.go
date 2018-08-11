// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"go.bmvs.io/iosrc"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	key := flag.String("key", "", "RestrictionsPasswordKey value")
	salt := flag.String("salt", "", "RestrictionsPasswordSalt value")

	flag.Parse()

	if *key == "" || *salt == "" {
		flag.Usage()
		return
	}

	pwd, err := iosrc.Crack(*key, *salt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pwd)
}
