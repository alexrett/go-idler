//+build darwin

package idler

// #cgo LDFLAGS: -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"

func (f *Idle) getIdleTime() float64 {
	return float64(C.CGEventSourceSecondsSinceLastEventType(C.kCGEventSourceStateHIDSystemState, C.kCGAnyInputEventType))
}
