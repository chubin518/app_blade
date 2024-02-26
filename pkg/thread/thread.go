package thread

import (
	"app_blade/pkg/logging"
	"fmt"
	"runtime"
)

// Start runs the given fn using another goroutine, recovers if fn panics.
func Start(fn func()) {
	go Run(fn)
}

// Run runs the given fn, recovers if fn panics.
func Run(fn func()) {
	defer Recover()
	fn()
}

// Recover used to dump stack info to file when catch panic
func Recover() {
	if r := recover(); r != nil {
		var msg string
		if err, ok := r.(error); ok {
			msg = err.Error()
		} else {
			msg = fmt.Sprintf("%#v", r)
		}
		buf := make([]byte, 1024)
		for {
			//the buf is no more than 64M, because Stack dumps no more than 64M
			n := runtime.Stack(buf, true)
			if n < len(buf) {
				buf = buf[:n] //trim unreadable characters
				break
			}
			buf = make([]byte, 2*len(buf))
		}

		logging.Default().Errorf("panic error: %s, stack: %s", msg, string(buf))
	}
}
