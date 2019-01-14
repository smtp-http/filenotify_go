package main

import (
	"github.com/smtp-http/filenotify_go/notify"
	"github.com/smtp-http/filenotify_go/conn"
)


func main() {
	monitor := notify.GetFileMonitorInstance()
	tcpserver := conn.GetServerInstance()
	go tcpserver.ServerRun("0.0.0.0","6688")
	monitor.SetTcpserver(tcpserver)
	monitor.Monitor()
}
