package main

import "hostname-chainer/server"

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
