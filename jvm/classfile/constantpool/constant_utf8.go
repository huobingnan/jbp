package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
	"unicode/utf16"
)

type Utf8Constant struct {
	value string
}

func (u *Utf8Constant) Tag() int {
	return ConstantUtf8Info
}

func (u *Utf8Constant) Value() interface{} {
	return u.value
}

func (u *Utf8Constant) String() string {
	return fmt.Sprintf("<CONSTANT_Utf8_info: %s>", u.value)
}

func (u *Utf8Constant) GoString() string {
	return u.String()
}

func NewUtf8Constant(reader *reader.ByteCodeReader) *Utf8Constant {
	length, _ := reader.ReadU2()
	utf8, _ := reader.ReadAny(int(length))
	return &Utf8Constant{
		value: decodeMUtf8(utf8),
	}
}

func decodeMUtf8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	charArrCount := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[charArrCount] = c
		charArrCount++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[charArrCount] = c
			charArrCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[charArrCount] = c&0x1F<<6 | char2&0x3F
			charArrCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[charArrCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charArrCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:charArrCount]
	runes := utf16.Decode(chararr)
	return string(runes)
}
