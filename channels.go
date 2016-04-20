package main

import (
  "fmt"
)

/*******************************************************************************
    Channels are a typed (meaning they support one data type) conduit through
    which you can send and receive values with channel operator "<-"
*******************************************************************************/


/*******************************************************************************
  The functions 'pinger' and 'ponger' take a string channel called 'c' that
                                    only accepts input
********************************************************************************/
func pinger(c chan<- string) {
  for i := 0; ; i++ {
    c <- "ping"           // infinitely sends "ping" to the channel c
  }
}

func ponger(c chan<- string) {
  for i := 0; ; i++ {
    c <- "pong"           // infinitely sends "pong" to the channel c
  }
}


/*******************************************************************************
  The functions 'pinger' and 'ponger' take a string channel called 'c' that
                          only accepts input
********************************************************************************/
func printer(c <-chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
