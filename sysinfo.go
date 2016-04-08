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
