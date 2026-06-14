package schema

import "github.com/quack-pot/quandry/internal/core"

type QueryBuilder struct {
}

func Query[T any](tx core.IQueryable) *QueryBuilder {
	return &QueryBuilder{}
}
