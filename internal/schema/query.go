package schema

import "github.com/quack-pot/quandry/internal/core"

type SqlExpression interface {
	Evaluate() string
}

type SqlPredicate[T any] struct {
	Expression SqlExpression
}

type QueryBuilder struct {
}

func Query[T any](tx core.IQueryable) *QueryBuilder {
	return &QueryBuilder{}
}
