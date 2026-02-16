package storage

type Condition struct {
	Field interface{}
	Value interface{}
	Op    Operation
}

type Operation string

const (
	OpEqual       Operation = "eq"
	OpNotEqual    Operation = "ne"
	OpGreaterThan Operation = "gt"
	OpLessThan    Operation = "lt"
	OpRange       Operation = "range"
	OpPrefix      Operation = "prefix"
	OpSuffix      Operation = "suffix"
	OpContains    Operation = "contains"
	OpIn          Operation = "in"
)
