package idler

// Dummy Idle struct
type Idle struct{}

// NewIdle creates a new instance of Idle
func NewIdle() *Idle {
	return &Idle{}
}

// GetIdleTime returns the number of seconds of inactivity of user activity
// IDLE time in seconds
func (f *Idle) GetIdleTime() int {
	if f == nil {
		return 0
	}
	return f.getIdleTime()
}
