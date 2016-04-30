package asyncproc

import "os/exec"

// AsyncProc is a pollable wrapper for a child process
type AsyncProc struct {
	Pid        int
	ExitStatus int
	Error      error
	done       bool
}

// MustStart launches a cmd returning AsyncProc or panicing if failure to launch
func MustStart(cmd *exec.Cmd) *AsyncProc {
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	return &AsyncProc{Pid: cmd.Process.Pid}
}

// Start launch as cmd and returns the error or an AsyncProc
func Start(cmd *exec.Cmd) (*AsyncProc, error) {
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return &AsyncProc{Pid: cmd.Process.Pid}, nil
}

// Exited updates status and
func (ap *AsyncProc) Exited() bool {
	if ap.done {
		return true
	}
	ap.done, ap.ExitStatus, ap.Error = pollExitStatus(ap.Pid)
	return ap.done
}
