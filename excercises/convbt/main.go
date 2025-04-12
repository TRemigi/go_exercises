package main

import (
	"fmt"
	"go_book/excercises/btz"
	"os"
	"strconv"
)

func main() {
  for _, arg := range os.Args[1:] {
    i, err := strconv.Atoi(arg)
    if err != nil {
      fmt.Fprintf(os.Stderr,"convbt: %v\n", err)
      os.Exit(1)
    }
    d := btz.Decimal(i)
    b := btz.DToB(d)
    o := btz.DToO(d)
    h := btz.DToH(d)
    fmt.Fprintf(os.Stdout, "Decimal: %d\nBinary: %d\nOctal: %d\nHexadecimal: %x", i, b, o, h)
    fmt.Fprintln(os.Stdout)
  }
}
