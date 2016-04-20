package main

import (
  "fmt"
  "time"
)

/*******************************************************************************
  The functions sendToC1 and sendToC2 infinitely send "From 1", and "From 2"
  to channel, sleeping a second each time.
********************************************************************************/
func sendToC1(channel chan string) {
    for {
        channel <- "From 1"
        time.Sleep(time.Second)
    }
}

func sendToC2(channel chan string) {
    for {
        channel <- "From 2"
        time.Sleep(time.Second)
    }
}


/*******************************************************************************
  This function takes 2 channels, and has a select statement, choosing to Print
  the value from each channel as soon as it's ready.
********************************************************************************/

func selectChannels(c1 chan string, c2 chan string) {
    for {
        select {
        case msg1 := <- c1:
            fmt.Println(msg1)
        case msg2 := <- c2:
            fmt.Println(msg2)
        }
    }
}

func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go sendToC1(c1);
  go sendToC2(c2);
  go selectChannels(c1,c2);

  var input string
  fmt.Scanln(&input)
}
