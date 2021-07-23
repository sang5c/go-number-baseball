package number

type Decision struct {
	result string
}

func (d Decision) GetResult() string {
	return d.result
}

var (
	nothing = Decision{"nothing"}
	ball    = Decision{"ball"}
	strike  = Decision{"strike"}
)
