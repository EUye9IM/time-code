package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/EUye9IM/tmcode"
)

func main() {
	ny := time.Now().Year()
	d := flag.String("d", "", "decode")
	b := flag.Int("b", ny, fmt.Sprintf("baseyear; default %v", ny))
	f := flag.String("f", "06/1/2.15", "time format; '' for unix timestamp")
	flag.Parse()
	if len(*d) > 0 {
		t, err := tmcode.Decode(*d, b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
			return
		}
		if *f == "" {
			fmt.Println(t.Unix())
		} else {
			fmt.Println(t.Format(*f))
		}
	} else {
		var t time.Time
		cmd := flag.Arg(0)
		if cmd == "" {
			t = time.Now()
		} else {
			var err error
			if *f == "" {
				ts, err := strconv.ParseInt(cmd, 10, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
					return
				}
				t = time.Unix(ts, 0)
			} else {
				t, err = time.Parse(*f, cmd)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
					return
				}
			}
		}

		fmt.Println(tmcode.Encode(t))
	}
}
