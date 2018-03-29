package main

import (
	"fmt"
	"github.com/emsyaiful/goblog/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebserver("6060")
}
