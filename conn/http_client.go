package conn

import (
	"errors"
)

type HttpClient struct {

}

func (c *TcpServer) HttpClient (msg string) bool {

	return true
}



func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	
	return a / b, nil
}