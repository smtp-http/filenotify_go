package conn

import (
	"fmt"
)

type TcpServer struct {

}

func (s *TcpServer) Notify (msg string) bool {
	fmt.Println(msg)
	return true
}