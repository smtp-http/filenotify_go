package main

import (
	"github.com/smtp-http/filenotify_go/notify"
)


func main() {
	var monitor *notify.FileMonitor = new(notify.FileMonitor)
	monitor.Monitor()
}