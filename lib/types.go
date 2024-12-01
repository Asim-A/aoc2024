package lib

type Output interface {
	Show()
}

type Execute interface {
	A(bool) int
	B(bool) int
}

type Challenge struct {
	Configuration DayConfig
}

type DayConfig struct {
	Day       int
	IsTest    bool
	Lines     []string
	TestLines []string
}
