package cypher

import (
	"runtime"
)

func CPUArch() string {
	return runtime.GOARCH
}