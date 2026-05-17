//go:build js && wasm

package main

import (
	"syscall/js"
	"time"

	"github.com/EUye9IM/tmcode"
)

func encode(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return "ERROR: missing unix timestamp"
	}
	ts := int64(args[0].Int())
	t := time.Unix(ts, 0).UTC()
	return tmcode.Encode(t)
}

func decode(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return "ERROR: missing code or baseYear"
	}
	code := args[0].String()
	baseYear := args[1].Int()
	t, err := tmcode.Decode(code, &baseYear)
	if err != nil {
		return "ERROR: " + err.Error()
	}
	return t.Unix()
}

func main() {
	js.Global().Set("tmcode_encode", js.FuncOf(encode))
	js.Global().Set("tmcode_decode", js.FuncOf(decode))
	select {}
}
