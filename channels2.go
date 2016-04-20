package main

import (
  "fmt"
)

/*******************************************************************************
Each value in the fibonacci sequence is calculating by adding the previous 2 values

        Here, n represents the number of values in the sequence.
  We start with 0 and 1, and send each value to the channel 'c' until we've
                                sent n terms
********************************************************************************/

func fibonacci(n int, c chan<- int) {
  x := 0
  y := 1

  for counter := 0; counter < n; counter++ {
    c <- x
    x = y
    y = x + y
  }
  close(c)  // Close tells the channel that no more values will be sent
}

func main() {
  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  for i := range c {
    fmt.Println(i)
  }
}
