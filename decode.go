package tmcode

import (
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"time"
)

// baseYear 为按年寻找最近可能的年份
// 默认按当前年份
func Decode(data string, baseYear *int) (time.Time, error) {
	var t time.Time
	if len(data) != 4 {
		return t, errors.New("invalid lenth")
	}
	b, err := base32.StdEncoding.DecodeString(data + "AAA=")
	if err != nil {
		return t, errors.New("invalid charactor")
	}
	dint := int(binary.BigEndian.Uint32(b) >> 12)
	if baseYear == nil {
		baseYear = new(int)
		*baseYear = time.Now().UTC().Year()
	}
	y250 := dint >> 12
	if y250 >= 250 || y250 < 0 {
		return t, errors.New("invalid data")
	}
	tmpy := (*baseYear - y250 + 125 - 1)
	y := (tmpy - tmpy%250 + y250)

	m := dint >> 8 & (1<<4 - 1)
	d := dint >> 3 & (1<<5 - 1)
	h := (dint & (1<<3 - 1)) * 3
	tt, err := time.Parse("2006/1/2/15", fmt.Sprintf("%v/%v/%v/%v", y, m, d, h))
	if err != nil {
		return t, errors.New("invalid data")
	}

	return tt, nil
}
