package schema

type Column interface {
	GenerateMeta() ColumnMeta
}

type IntColumn[C Constraint] struct{}
