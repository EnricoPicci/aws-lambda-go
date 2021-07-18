package main

import (
	"fmt"

	"github.com/serverless/helloYou/src/server"
)

func main() {
	fmt.Println("Command to start the http server")

	server.StartHttpSrv()
}
