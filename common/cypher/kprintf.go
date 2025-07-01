package cypher

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Verbosity int
	LogName   string
	logger    *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

func ReopenLogs() {
	// در Go معمولاً نیازی به reopen نیست، مگر اینکه فایل لاگ عوض شود
}

func HexDump(data []byte) {
	fmt.Printf("%x\n", data)
}

func KWrite(fd *os.File, buf []byte) (int, error) {
	return fd.Write(buf)
}

func KPrintf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func VKPrintf(verbosityLevel int, format string, v ...interface{}) {
	if verbosityLevel > Verbosity {
		return
	}
	KPrintf(format, v...)
}