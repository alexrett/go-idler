package idler

type Idle struct{}

// GetIdleTime returns the number of seconds of inactivity of user activity
// IDLE time in seconds
func (f *Idle) GetIdleTime() int {
	if f == nil {
		return 0
	}
	return f.getIdleTime()
}
