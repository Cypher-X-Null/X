package cypher

import (
	"os"
	"syscall"
)

type ProcStats struct {
	PID  int
	PPID int
	// سایر فیلدها را بر اساس نیاز اضافه کنید
}

func ReadProcStats(pid int) (*ProcStats, error) {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	// اطلاعات بیشتر را با syscall یا خواندن /proc بدست آورید
	return &ProcStats{PID: proc.Pid}, nil
}