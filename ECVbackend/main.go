package main

import "ecvbackend/setup"

func main() {
	router := setup.SetupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}
