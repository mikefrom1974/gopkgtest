package main

import (
	"fmt"

	"github.com/mikefrom1974/gopkgtest/server"
)

func main() {
	s := server.New("testServer")
	fmt.Println(s.Name)
	fmt.Println(s.Process())
}
