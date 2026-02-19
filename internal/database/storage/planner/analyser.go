package planner

type Analysis struct {
	Cost uint32
	Rows uint32
}

func NewAnalysis(cost uint32, rows uint32) Analysis {
	return Analysis{
		Cost: cost,
		Rows: rows,
	}
}
