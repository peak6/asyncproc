package asyncproc

import (
	"os/exec"
	"testing"
	"time"
)

func TestExitCode(t *testing.T) {
	// proably not windows
	ap := MustStart(exec.Command("bash", "-c", "exit 2"))
	waitForExit(t, time.Second, ap)
	if ap.ExitStatus != 2 {
		t.Fatal("Expected exit code 2, got:", ap.ExitStatus)
	}
}

func TestExited(t *testing.T) {
	waitForExit(t, time.Second, sleep(t, ".01"))
}

func TestStillRunning(t *testing.T) {
	ap := sleep(t, "1")
	if ap.Exited() {
		t.Error("Expected process to still be running")
	}
}

func sleep(t *testing.T, sec string) *AsyncProc {
	return MustStart(exec.Command("sleep", sec)) // not sure if windows has this
}

func waitForExit(t *testing.T, maxWait time.Duration, ap *AsyncProc) {
	for die := time.Now().Add(maxWait); time.Now().Before(die); {
		if ap.Exited() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("Process did not exit within:", maxWait)
}
