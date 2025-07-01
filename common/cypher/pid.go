package cypher

import (
	"os"
	"time"
)

type ProcessID struct {
	IP    string
	Port  int
	PID   int
	UTime int64
}

func NewProcessID(ip string, port int) ProcessID {
	return ProcessID{
		IP:    ip,
		Port:  port,
		PID:   os.Getpid(),
		UTime: time.Now().Unix(),
	}
}