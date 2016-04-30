// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package asyncproc

import (
	"fmt"
	"syscall"
)

func pollExitStatus(pid int) (bool, int, error) {
	var status syscall.WaitStatus
	// var rusage syscall.Rusage
	opts := syscall.WNOHANG
	p1, err := syscall.Wait4(pid, &status, opts, nil)
	if err != nil {
		return true, -1, err
	} else if p1 == pid {
		return true, status.ExitStatus(), nil
	} else if p1 == -1 {
		panic("Shoud never get a pid of -1 with no error")
	} else if p1 == 0 {
		// Still Running
		return false, 0, nil
	} else {
		panic(fmt.Sprintf("Received unexpected pid: %d", p1))
	}
}
