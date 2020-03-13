//+build windows

package idler

// #include <windows.h>
import "C"
import "unsafe"
import "syscall"

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	getLastInputInfo = user32.MustFindProc("GetLastInputInfo")
	getTickCount     = kernel32.MustFindProc("GetTickCount")
	lastInputInfo    struct {
		cbSize uint32
		dwTime uint32
	}
)

func (f *Idle) getIdleTime() float64 {
	lastInputInfo.cbSize = uint32(unsafe.Sizeof(lastInputInfo))
	currentTickCount, _, _ := getTickCount.Call()
	r1, _, _ := getLastInputInfo.Call(uintptr(unsafe.Pointer(&lastInputInfo)))
	if r1 == 0 {
		return 0
	}
	return float64(uint32(uint32(currentTickCount)-lastInputInfo.dwTime) / 1000)
}
