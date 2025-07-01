package cypher

import (
	"time"
)

func GetUTime() int64 {
	return time.Now().UnixNano()
}

func GetDoubleTime() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}