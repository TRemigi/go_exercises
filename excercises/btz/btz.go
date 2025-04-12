// Package btz converts a given decimal number to Binary and Octal.
package btz

import (
	"fmt"
)

type Decimal int
type Binary int
type Octal int
type Hexadecimal []byte

const (
	DecimalBase     int = 10
	BinaryBase      int = 2
	OctalBase       int = 8
	HexadecimalBase int = 16
)

var HexMap = map[Decimal]byte{
	0:  0x0,
	1:  0x1,
	2:  0x2,
	3:  0x3,
	4:  0x4,
	5:  0x5,
	6:  0x6,
	7:  0x7,
	8:  0x8,
	9:  0x9,
	10: 0xA,
	11: 0xB,
	12: 0xC,
	13: 0xD,
	14: 0xE,
	15: 0xF,
}

func (d Decimal) ToHexByte() byte { return HexMap[d] }
func (d Decimal) String() string  { return fmt.Sprintf("%d base %d", d, DecimalBase) }
func (b Binary) String() string   { return fmt.Sprintf("%d base %d", b, BinaryBase) }
func (o Octal) String() string    { return fmt.Sprintf("%d base %d", o, OctalBase) }
