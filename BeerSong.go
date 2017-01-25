package main

import "fmt"
import "strconv"

type bottle interface {
  Ninety_Nine_Bottles_of_Beer(n int, b beers) int
}

type beers struct {
  count int
}

func (brs beers) Ninety_Nine_Bottles_of_Beer(n int, b beers) int {
}

func main() {
}
