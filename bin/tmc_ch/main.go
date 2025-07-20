package main

import (
	"encoding/base32"
	"flag"
	"fmt"

	"github.com/EUye9IM/tmcode"
)

func main() {
	flag.Parse()
	s := flag.Arg(0)
	b, _ := base32.StdEncoding.DecodeString(s + tmcode.Padding)
	fmt.Println(base32.NewEncoding(tmcode.CodeStr).EncodeToString(b)[:4])
}
