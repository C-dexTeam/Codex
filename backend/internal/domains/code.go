package domains

type CodeResponse struct {
	Correct        bool
	Output         string
	BuildError     string
	Err            string
	CorrectTestsID []string
	WrongTestID    string
}
