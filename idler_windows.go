//+build windows

package idler

// #include <windows.h>
import "C"

// todo Нужно проверить, ибо винды нет, возможно так это не работает;
// Для подсказки смотри мое прошлое решение - https://github.com/Allexin/TrackYourTime/issues/66#issue-142088869
func (f *Idle) getIdleTime() float64 {
	var li C.LASTINPUTINFO
	li.cbSize = C.sizeof(C.LASTINPUTINFO)
	C.GetLastInputInfo(&li)
	var te C.DWORD = C.GetTickCount()
	return float64((te - li.dwTime) / 1000)
}
