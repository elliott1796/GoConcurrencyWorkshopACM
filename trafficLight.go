package main

import (
  "fmt"
  "time"
)

/***********************************************************************
 drivingNorth and turningLeft both infinitely send their rude gestures
 to the the channel that was passed as a function parameter
*************************************************************************/
func drivingNorth(greenLight chan<- string) {
  for {
    greenLight <- "BEEP BEEP, GOING NORTH PAL, OUTTA MY WAY!!!"
    time.Sleep(time.Second *2)
  }
}

func turningLeft(greenLight chan<- string) {
  for {
    greenLight <- "DON'T CALL ME PAL BUDDY, I'M TURNING LEFT!"
    time.Sleep(time.Second)
  }
}



func trafficIntersection(light1 chan string, light2 chan string) {
  for {
    select {
    case northRudeGesture := <- light1:
      fmt.Println(northRudeGesture)
    case leftRudeGesture := <- light2:
      fmt.Println(leftRudeGesture)
    }
  }
}

func main() {
  northGreenLight := make(chan string)
  leftGreenLight := make(chan string)

  go drivingNorth(northGreenLight)
  go turningLeft(leftGreenLight)
  go trafficIntersection(northGreenLight,leftGreenLight)


  // Ignore
  var input string
  fmt.Scanln(&input)
}
