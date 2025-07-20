package tmcode

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"time"
)

func Encode(t time.Time) string {
	t = t.UTC()
	y, m, d := t.Date()
	h := t.Hour()
	data := uint32((y%250)<<12 | int(m)<<8 | d<<3 | (h / 3))
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data<<12)
	return base32.NewEncoding(CodeStr).EncodeToString(bytebuf.Bytes())[:4]
}
