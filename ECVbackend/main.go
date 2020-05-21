package main

import (
	"ecvbackend/middleware"
	"ecvbackend/setup"
	"os"
)

func main() {
	setup.Init()
	router := setup.SetupRouter()
	middleware.ListenPort = os.Args[1]
	router.Run(":" + os.Args[1])
}
