package server

import (
	"github.com/mikefrom1974/gopkgtest/data"
)

type MyServer struct {
	Name string
	Data *data.MyData
}

func New(name string) *MyServer {
	s := new(MyServer)
	s.Name = name
	s.Data = data.New("Hello, World!")
	return s
}

func (s *MyServer) Process() string {
	return s.Data.TestString
}
