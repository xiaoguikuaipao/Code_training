package intsToString

import (
	"bytes"
	"fmt"
)

func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteRune(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteString("]")
	return buf.String()
}

//无论是写Byte Rune String 最终都是转成一个字节
