package number

type Decision string

func (d Decision) IsStrike() bool {
	return d == "strike"
}

func (d Decision) IsBall() bool {
	return d == "ball"
}

const (
	nothing = Decision("nothing")
	ball    = Decision("ball")
	strike  = Decision("strike")
)
