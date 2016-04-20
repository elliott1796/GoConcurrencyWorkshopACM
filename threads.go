package main

import (
  "fmt"
  "time"
  "math/rand"
)

/**********************************************************
    Process is a function that takes an integer,
    It iterates counter from 1 through 3, printing counter
    and the integer passed to it, sleeping a random amount
                of milliseconds after each time.
***********************************************************/
func process(numberThread int) {
  for counter := 1; counter <= 3; counter++ {
    go fmt.Println("Thread:", numberThread, ":", counter)
    timeAmount := time.Duration(rand.Intn(50))
    time.Sleep(time.Millisecond * timeAmount)
  }
}

/********************************************************
    In main we're again iterating from 1 through 3,
    each time starting off a thread that executes the
    process function.
*********************************************************/
func main() {
  for counter := 1; counter <= 3; counter++ {
    go process(counter)
  }

  // next two lines are just to be sneaky
  var input string
  fmt.Scanln(&input)
}
