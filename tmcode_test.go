package tmcode_test

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"testing"
	"time"

	"github.com/EUye9IM/tmcode"
)

func TestEncodingSucc(t *testing.T) {
	table := [][3]string{
		{"2021-07-29 14:15:46", "CV7M", "2021-07-29 12:00:00"},
		{"2018-08-19 08:19:37", "CKE2", "2018-08-19 06:00:00"},
		{"2018-01-23 17:31:41", "CIN5", "2018-01-23 15:00:00"},
		{"2017-09-08 12:49:15", "CGKE", "2017-09-08 12:00:00"},
		{"2016-11-13 19:02:58", "CC3O", "2016-11-13 18:00:00"},
		{"2015-09-17 06:54:31", "B6MK", "2015-09-17 06:00:00"},
		{"2015-05-30 19:23:02", "B5PW", "2015-05-30 18:00:00"},
		{"2015-04-25 14:11:26", "B5GM", "2015-04-25 12:00:00"},
		{"2014-04-02 07:46:47", "BZAS", "2264-04-02 06:00:00"},
		{"2013-05-24 13:44:49", "BVOE", "2263-05-24 12:00:00"},
		{"2012-04-11 18:43:12", "BRC6", "2262-04-11 18:00:00"},
		{"2012-04-11 16:38:36", "BRC5", "2262-04-11 15:00:00"},
	}
	base := 2140
	tf := "2006-01-02 15:04:05"

	for _, i := range table {
		tm, err := time.Parse(tf, i[0])
		if err != nil {
			t.Fatal(i[0], err)
		}
		s := tmcode.Encode(tm)
		result, err := tmcode.Decode(s, &base)
		if err != nil {
			t.Fatal(i[0], err)
		}
		resultS := result.Format(tf)
		if s != i[1] || resultS != i[2] {
			t.Errorf("fail: %v -> %v -> %v", i[1], s, resultS)
		}
	}
}

func TestDecodingFail(t *testing.T) {
	fun := func(y, m, d, h int) string {
		data := uint32(y<<12 | int(m)<<8 | d<<3 | h)
		bytebuf := bytes.NewBuffer([]byte{})
		binary.Write(bytebuf, binary.BigEndian, data<<12)
		return base32.StdEncoding.EncodeToString(bytebuf.Bytes())[:4]
	}
	table := [][2]string{
		{"AAAAA", "invalid lenth"},
		{"AAA8", "invalid charactor"},
		{fun(250, 1, 1, 0), "invalid data"},
		{fun(-1, 1, 1, 0), "invalid data"},
		{fun(0, 0, 1, 0), "invalid data"},
		{fun(0, 13, 1, 0), "invalid data"},
		{fun(0, 1, 0, 0), "invalid data"},
		{fun(0, 1, 32, 0), "invalid data"},
	}
	for i, v := range table {
		if _, err := tmcode.Decode(v[0], nil); err == nil || err.Error() != v[1] {
			t.Errorf("fail: %v (%v) -> %v", v[0], i, err)
		}
	}
}
