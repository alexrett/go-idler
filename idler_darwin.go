//+build darwin

package idler

// #cgo LDFLAGS: -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"

func (f *Idle) getIdleTime() int {
	return int(C.CGEventSourceSecondsSinceLastEventType(C.kCGEventSourceStateHIDSystemState, C.kCGAnyInputEventType))
}
