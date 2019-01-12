package conn

import (
	"fmt"
	"github.com/smtp-http/zero"
	"container/list"
	"sync"
)


type TcpServer struct {
	SessionList *list.List
}

func (s *TcpServer) Notify (msg []byte) bool {

	for e := s.SessionList.Front(); e != nil; e = e.Next() {
		session := e.Value.(*zero.Session)
		conn := session.GetConn()
		
		conn.SendBytes(msg)
	}

	return true
}

 
var instance *TcpServer
var once sync.Once
 
func GetServerInstance() *TcpServer {
    once.Do(func() {
        instance = &TcpServer{}
    })
    return instance
}





func (s *TcpServer) ServerRun(ip string,port string) {
	
	s.SessionList = list.New()

	ss, err := zero.NewSocketService(ip + ":" + port)
	if err != nil {
		return
	}

	ss.RegMessageHandler(HandleMessage)
	ss.RegConnectHandler(HandleConnect)
	ss.RegDisconnectHandler(HandleDisconnect)

	ss.Serv()
}


func HandleMessage(s *zero.Session, msg *zero.Message) {
	fmt.Println("receive msgID:", msg)
	fmt.Println("receive data:", string(msg.GetData()))
}

func HandleDisconnect(s *zero.Session, err error) {
	fmt.Println(s.GetConn().GetName() + " lost.")
}

func HandleConnect(s *zero.Session) {
	fmt.Println(s.GetConn().GetName() + " connected." )
	tcpserver := GetServerInstance()
	tcpserver.SessionList.PushBack(s)
}