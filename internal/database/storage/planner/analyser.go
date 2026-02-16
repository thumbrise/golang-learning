package planner

type Analysis struct {
	Cost float64
	Rows int
}

func NewAnalysis(cost float64, rows int) Analysis {
	return Analysis{
		Cost: cost,
		Rows: rows,
	}
}
