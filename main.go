package main

import (
	"github.com/Abdulrasheed1729/hng12-stage1/api"
)

func main() {

	port := ":8080"

	server := api.NewAPIServer(port)

	server.Run()

}
