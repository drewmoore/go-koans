package go_koans

import "fmt"

func divideFourBy(i int) int {
  defer func() {
    if e := recover(); e != nil {
      fmt.Println("Sorry, yo, can't divide by", e)
    }
  }()
  if i == 0 { panic(i) }
  return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
  assert(divideFourBy(__divisor__) == 0) // panics are exceptional errors at runtime

  n := divideFourBy(2)
  assert(n == 2) // panics are exceptional errors at runtime
}
