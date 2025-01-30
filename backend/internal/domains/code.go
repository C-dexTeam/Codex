package domains

type CodeResponse struct {
	Correct        bool
	Output         string
	BuildError     string
	Err            error
	CorrectTestsID []string
	WrongTestID    string
}
