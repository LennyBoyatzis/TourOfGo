package main

import (
  "fmt"
  "flag"
  "os"
  "chat-app/lib"
)

func main() {
  fmt.Println("Hello")
  var isHost bool
  flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
  flag.Parse()

  if isHost {
    connIP := os.Args[2]
    lib.RunHost(connIP)
    fmt.Println("is host")
  } else {
    connIP := os.Args[1]
    lib.RunHost(connIP)
    fmt.Println("is guest")
  }
}
