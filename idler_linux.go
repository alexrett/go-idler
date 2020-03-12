//+build linux

package idler

// #cgo LDFLAGS: -lXss -lX11
// #include <X11/extensions/scrnsaver.h>
import "C"

func (f *Idle) getIdleTime() float64 {
	var info *C.XScreenSaverInfo
	var display *C.Display = C.XOpenDisplay(C.CString(""))
	info = C.XScreenSaverAllocInfo()
	defaultRootWindow := C.XDefaultRootWindow(display)
	if int(defaultRootWindow) == -1 {
		return 0
	}
	C.XScreenSaverQueryInfo(display, C.Drawable(defaultRootWindow), info)

	var idleTime uint32
	idleTime = uint32(info.idle)

	return float64(idleTime / 1000)
}
