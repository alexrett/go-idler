//+build linux

package idler

// #cgo LDFLAGS: -lXss -lX11
// #include <X11/extensions/scrnsaver.h>
import "C"

func (f *Idle) getIdleTime() int {
	var info *C.XScreenSaverInfo = C.XScreenSaverAllocInfo()
	var display *C.Display = C.XOpenDisplay(C.CString(""))

	defaultRootWindow := C.XDefaultRootWindow(display)
	if int(defaultRootWindow) == -1 {
		return 0
	}

	C.XScreenSaverQueryInfo(display, C.Drawable(defaultRootWindow), info)
	
	C.XCloseDisplay(display)
	
	return int(uint32(info.idle) / 1000)
}
