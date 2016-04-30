# asyncproc
Asynchronous Process Monitor

Simple library to allow polling of process exit status without blocking on a syscall.

## Motivation  
We manage a system comprised of hundreds of executables running on a single machine.  

Go's standard way of process exit monitoring is to do a blocking call to wait4 (on unix hosts).
This causes the go runtime to spawn an addition thread per blocking call.  

These threads, while sleeping, add a fair amount of memory overhead to the waiting process.

This library allows us to poll for exited processes rather than block.

TODO: Port to windows

Example usage:

```go
package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/peak6/asyncproc"
)

func main() {
	// Create a wrapped child process or panic
	ap := asyncproc.MustStart(exec.Command("bash", "-c", "sleep 1;exit 2"))
	for {
		if ap.Exited() {
			fmt.Println("Exited, with status:", ap.ExitStatus)
			break
		}
		fmt.Println("Still waiting")
		time.Sleep(250 * time.Millisecond)
	}
}
```
