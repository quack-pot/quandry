package schema

import "github.com/quack-pot/quandry/internal/core"

type IColumnType interface {
	SqlType(dialect core.ISqlDialect) string
	ApplyConstraints(meta *ColumnMeta)
}

// Integer Columns ------------------------------------------------------------

type IntColumn[C IConstraint] struct{}

func (c IntColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.IntType() }
func (c IntColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type SmallIntColumn[C IConstraint] struct{}

func (c SmallIntColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.SmallIntType() }
func (c SmallIntColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type BigIntColumn[C IConstraint] struct{}

func (c BigIntColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.BigIntType() }
func (c BigIntColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// Decimal Columns ------------------------------------------------------------

type FloatColumn[C IConstraint] struct{}

func (c FloatColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.FloatType() }
func (c FloatColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type DoubleColumn[C IConstraint] struct{}

func (c DoubleColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.DoubleType() }
func (c DoubleColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// Monetary Columns -----------------------------------------------------------

type MoneyColumn[C IConstraint] struct{}

func (c MoneyColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.MoneyType() }
func (c MoneyColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// Boolean Columns ------------------------------------------------------------

type BooleanColumn[C IConstraint] struct{}

func (c BooleanColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.BooleanType() }
func (c BooleanColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// Character Columns ----------------------------------------------------------

type TextColumn[C IConstraint] struct{}

func (c TextColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.TextType() }
func (c TextColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// Time Columns ---------------------------------------------------------------

type DateColumn[C IConstraint] struct{}

func (c DateColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.DateType() }
func (c DateColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type TimeColumn[C IConstraint] struct{}

func (c TimeColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.TimeType() }
func (c TimeColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type DateTimeColumn[C IConstraint] struct{}

func (c DateTimeColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.DateTimeType() }
func (c DateTimeColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type DateTimeZoneColumn[C IConstraint] struct{}

func (c DateTimeZoneColumn[C]) SqlType(dialect core.ISqlDialect) string {
	return dialect.DateTimeZoneType()
}
func (c DateTimeZoneColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type IntervalColumn[C IConstraint] struct{}

func (c IntervalColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.IntervalType() }
func (c IntervalColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// Json Columns ---------------------------------------------------------------

type JsonColumn[C IConstraint] struct{}

func (c JsonColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.JsonType() }
func (c JsonColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type JsonBinaryColumn[C IConstraint] struct{}

func (c JsonBinaryColumn[C]) SqlType(dialect core.ISqlDialect) string {
	return dialect.JsonBinaryType()
}
func (c JsonBinaryColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// Network Columns ------------------------------------------------------------

type CidrColumn[C IConstraint] struct{}

func (c CidrColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.CidrType() }
func (c CidrColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type INetColumn[C IConstraint] struct{}

func (c INetColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.INetType() }
func (c INetColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type MacAddress6Column[C IConstraint] struct{}

func (c MacAddress6Column[C]) SqlType(dialect core.ISqlDialect) string {
	return dialect.MacAddress6Type()
}
func (c MacAddress6Column[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

type MacAddress8Column[C IConstraint] struct{}

func (c MacAddress8Column[C]) SqlType(dialect core.ISqlDialect) string {
	return dialect.MacAddress8Type()
}
func (c MacAddress8Column[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

// UUID Columns ---------------------------------------------------------------

type UuidColumn[C IConstraint] struct{}

func (c UuidColumn[C]) SqlType(dialect core.ISqlDialect) string { return dialect.UuidType() }
func (c UuidColumn[C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}
