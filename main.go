package main

import "recursive-echo/server"

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
