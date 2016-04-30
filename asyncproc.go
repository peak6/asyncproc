package asyncproc

import (
	"errors"
	"os/exec"
)

// ErrNotStarted is returned by Wrap when you pass and unstarted cmd
var ErrNotStarted = errors.New("process not started")

// AsyncProc is a pollable wrapper for a child process
type Proc struct {
	Pid        int
	ExitStatus int
	Error      error
	done       bool
}

// Wrap takes a running cmd and returns an AsyncProc
func Wrap(cmd *exec.Cmd) (*Proc, error) {
	if cmd.Process == nil {
		return nil, ErrNotStarted
	}
	return &Proc{Pid: cmd.Process.Pid}, nil
}

// MustStart launches a cmd returning AsyncProc or panicing if failure to launch
func MustStart(cmd *exec.Cmd) *Proc {
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	return &Proc{Pid: cmd.Process.Pid}
}

// Start launch as cmd and returns the error or an AsyncProc
func Start(cmd *exec.Cmd) (*Proc, error) {
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return &Proc{Pid: cmd.Process.Pid}, nil
}

// Exited updates status and
func (ap *Proc) Exited() bool {
	if ap.done {
		return true
	}
	ap.done, ap.ExitStatus, ap.Error = pollExitStatus(ap.Pid)
	return ap.done
}
