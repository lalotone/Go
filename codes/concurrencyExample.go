package main

import (
  "fmt"
)

func ponger(z chan string) {
  for i := 0; ; i++ {
    z <- "pong"
  }
}
func pinger(c chan string) {
  for i := 0; ; i++ {
    c <- "ping"
  }
}
func penger(x chan string) {
  for i := 0; ; i++ {
    x <- "peng"
  }
}
func printer(c,z,x chan string) {
  for {
    msgThree := <- x
    msgTwo := <- z
    msg := <- c
    fmt.Println(msg)
    fmt.Println(msgTwo)
    fmt.Println(msgThree)
  }
}

func main() {
  var c chan string = make(chan string)
  var z chan string = make(chan string)
  var x chan string = make(chan string)
  go pinger(c)
  go ponger(z)
  go penger(x)
  go printer(c,z,x)

  var input string
  fmt.Scanln(&input)
}
