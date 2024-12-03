package lib

type Output interface {
	Show()
}

type Execute interface {
	A() int
	B() int
}

type Challenge struct {
	Arguments Arguments
}
