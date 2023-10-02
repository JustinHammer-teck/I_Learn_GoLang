package server

import (
	"fmt"
	"net"
	"os"

	"github.com/JustinHammer-teck/I_Learn_GoLang/TcpServer/types"
)


func NewTcpServer(config types.ServerConfig){
  connection, err := net.Listen("tcp", config.IP + ":" + config.PORT)
  if err != nil {
    fmt.Println("Something went wrong", err)
    os.Exit(1)
  }

  host, port, err := net.SplitHostPort(connection.Addr().String())
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("Listen to port %s at the ip %s\n", port, host)

  defer connection.Close()
  for {
    conn, err := connection.Accept()
    if err != nil {
      fmt.Println("Cannot accept tcp connection", err)
    }
    go Handle(conn)
  }
}

func Handle(conn net.Conn){
  buffer := make([]byte, 1024)
  _, err := conn.Read(buffer)
  if err != nil {
    fmt.Println(err)

    return
  }

  fmt.Println("Received: ", string(buffer))

  conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

  conn.Close()
}