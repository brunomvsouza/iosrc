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
