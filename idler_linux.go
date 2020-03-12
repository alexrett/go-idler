//+build linux

package idler

// #cgo LDFLAGS: -lXss -lX11
// #include <X11/extensions/scrnsaver.h>
import "C"
import "fmt"

func (f *Idle) getIdleTime() float64 {
	var info *C.XScreenSaverInfo
	var display *C.Display

	defaultRootWindow := C.XDefaultRootWindow(display)

	C.XScreenSaverQueryInfo(display, C.Drawable(defaultRootWindow), info)

	var idleTime uint32
	idleTime = uint32(info.idle)

	return float64(idleTime)
}
