package domain

//["counter","left","right","xmid","fXmid","error"]
type BisectionLine struct {
	Counter string
	Left    string
	Right   string
	Xmid    string
	Fxmid   string
	Error   string
}

type Bisection struct {
	Lines []BisectionLine
}
