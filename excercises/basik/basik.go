// Package basik converts a given decimal number to binary, octal, and hexadecimal.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
  for _, d := range os.Args[1:] {
    i, err := strconv.Atoi(d)
    if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
      os.Exit(1)
    }
    b := strconv.FormatInt(int64(i), 2)
    o := strconv.FormatInt(int64(i), 8)
    h := strconv.FormatInt(int64(i), 16)
    fmt.Printf("decimal: %s\nbinary: %s\noctal: %s\nhexadecimal: %s\n", d, b, o, h)
  }
}
