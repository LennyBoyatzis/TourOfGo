package lib

import (
  "bufio"
  "net"
  "log"
  "fmt"
)

const port = "8080"

func RunHost(ip string) {
  ipAndPort := ip + ":" + port
  listener, listenErr := net.Listen("tcp", ipAndPort)
  if listenErr != nil {
    log.Fatal("Error: ", listenErr)
  }

  conn, acceptErr := listener.Accept()
  if acceptErr != nil {
    log.Fatal("Error: ", acceptErr)
  }

  reader := bufio.NewReader(conn)
  message, readErr := reader.ReadString('\n')
  if readErr != nil {
    log.Fatal("Error:", readErr)
  }
  fmt.Println("Message received: ", message)
}

func RunGuest(ip string) {

}
