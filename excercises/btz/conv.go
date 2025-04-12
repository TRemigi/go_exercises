package btz

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// DToB converts a decimal number to its binary representation.
func DToB(d Decimal) Binary {
	i := int(d)
	if i < 1 {
		return Binary(0)
	}

	var b []int
	for r := 0; i > 0; {
		i, r = divEx(i, BinaryBase)
		b = append([]int{r}, b...)
	}

	return Binary(intSliceToInt(b))
}

// DToO converts a decimal number to its octal representation.
func DToO(d Decimal) Octal {
	i := int(d)
	if i < 1 {
		return Octal(0)
	}

	var o []int
	for r := 0; i > 0; {
		i, r = divEx(i, OctalBase)
		o = append([]int{r}, o...)
	}

	return Octal(intSliceToInt(o))
}

// DToH converts a decimal number to its Hexadecimal representation.
func DToH(d Decimal) Hexadecimal {
	i := int(d)
	if i < 1 {
		return Hexadecimal([]byte{0x0})
	}

	var b []byte
	for r := 0; i > 0; {
		i, r = divEx(i, HexadecimalBase)
		h := Decimal(r).ToHexByte()
    fmt.Fprintf(os.Stdout, "%x", h)
		b = append([]byte{h}, b...)
	}
  fmt.Printf("%x", b)

	return Hexadecimal(b)
}

// devEx returns both the quotient and remainder for a given numerator and denominator.
func divEx(n int, d int) (int, int) {
	return n / d, n % d
}

// intSliceToInt converts a given []int to an int.
func intSliceToInt(s []int) int {
	var sb strings.Builder
	for _, i := range s {
		sb.WriteString(strconv.Itoa(i))
	}
	i, err := strconv.Atoi(sb.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "btz: %v\n", err)
		os.Exit(1)
	}
	return i
}
