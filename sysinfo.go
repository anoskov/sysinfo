package sysinfo

type Uptime struct {
	Duration float64
}

type AverageLoad struct {
	One, Five, Fifteen float64
}

type CPU struct {
	Count int
}

type Goroutine struct {
	Count int
}

type RAM struct {
	Total      uint64
	Used       uint64
	Free       uint64
	ActualFree uint64
	ActualUsed uint64
}
