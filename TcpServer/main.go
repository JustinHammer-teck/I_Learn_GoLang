package main

import (
	"flag"

	"github.com/JustinHammer-teck/I_Learn_GoLang/TcpServer/server"
	"github.com/JustinHammer-teck/I_Learn_GoLang/TcpServer/types"
)

var (
  portConfig = flag.String("port", "9051", "Custom port for server")
  ipConfig   = flag.String("ip", "localhost", "Custom Ip address for tcp server")
)

func main(){

  flag.Parse()

  serverConfig := types.ServerConfig{IP : *ipConfig, PORT: *portConfig};
  server.NewTcpServer(serverConfig)
}