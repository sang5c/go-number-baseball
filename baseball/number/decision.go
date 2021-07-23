package number

type Decision struct {
	result string
}

func (d Decision) IsStrike() bool {
	return d.result == "Strike"
}

func (d Decision) IsBall() bool {
	return d.result == "Ball"
}

var (
	Nothing = Decision{"Nothing"}
	Ball    = Decision{"Ball"}
	Strike  = Decision{"Strike"}
)
