package idler

type Idle struct{}

func (f *Idle) GetIdleTime() float64 {
	if f == nil {
		return 0
	}
	return f.getIdleTime()
}
