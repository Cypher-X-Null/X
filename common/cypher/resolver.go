package cypher

import (
	"net"
	"os"
)

func LookupHost(name string) ([]string, error) {
	return net.LookupHost(name)
}

func DetectHostname() (string, error) {
	return os.Hostname()
}