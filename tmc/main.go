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
	d := flag.Bool("d", false, "decode")
	b := flag.Int("b", ny, fmt.Sprintf("baseyear; default %v", ny))
	f := flag.String("f", "", "time format; use Unix time defaultly")
	flag.Parse()
	cmd := fmt.Sprint(flag.Args())
	cmd = cmd[1 : len(cmd)-1]
	if *d {
		t, err := tmcode.Decode(cmd, b)
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
