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
  if n == 1 {
    fmt.Println(strconv.Itoa(n) + " bottle of beer on the wall " + strconv.Itoa(n) + " bottle of beer take one down pass it around no more bottles of beer on the wall")
    return 0
  } else if n == 2 {
    fmt.Println(strconv.Itoa(n) + " bottles of beer on the wall " + strconv.Itoa(n) + " bottles of beer take one down pass it around " + strconv.Itoa(n - 1) + " bottles of beer on the wall")
    return b.Ninety_Nine_Bottles_of_Beer(n-1, b)
  }
  fmt.Println(strconv.Itoa(n) + " bottles of beer on the wall " + strconv.Itoa(n) + " bottles of beer take one down pass it around " + strconv.Itoa(n - 1) + " bottles of beer on the wall")
  return b.Ninety_Nine_Bottles_of_Beer(n-1, b)
}

func main() {
  brs := beers{ count: 99 }
  count := brs.count
  brs.Ninety_Nine_Bottles_of_Beer(count, brs)
}
