package schema

import (
	"github.com/quack-pot/quandry/internal/core"
)

type IColumnType interface {
	SqlType(dialect core.ISqlDialect) string
	ApplyConstraints(meta *ColumnMeta)
}

type ColumnType[T any, C IConstraint] struct {
	IColumnType
}

func (ColumnType[T, C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

func (ColumnType[T, C]) SqlType(dialect core.ISqlDialect) string {
	return SqlTypeStringFor[T](dialect)
}

// Integer Columns ------------------------------------------------------------

type IntColumn[C IConstraint] = ColumnType[SqlInt, C]
type SmallIntColumn[C IConstraint] = ColumnType[SqlSmallInt, C]
type BigIntColumn[C IConstraint] = ColumnType[SqlBigInt, C]

// Decimal Columns ------------------------------------------------------------

type FloatColumn[C IConstraint] = ColumnType[SqlFloat, C]
type DoubleColumn[C IConstraint] = ColumnType[SqlDouble, C]

// Monetary Columns -----------------------------------------------------------

type MoneyColumn[C IConstraint] = ColumnType[SqlMoney, C]

// Boolean Columns ------------------------------------------------------------

type BooleanColumn[C IConstraint] = ColumnType[SqlBoolean, C]

// Character Columns ----------------------------------------------------------

type TextColumn[C IConstraint] = ColumnType[SqlText, C]

// Time Columns ---------------------------------------------------------------

type DateColumn[C IConstraint] = ColumnType[SqlDate, C]
type TimeColumn[C IConstraint] = ColumnType[SqlTime, C]
type DateTimeColumn[C IConstraint] = ColumnType[SqlDateTime, C]
type DateTimeZoneColumn[C IConstraint] = ColumnType[SqlDateTimeZone, C]
type IntervalColumn[C IConstraint] = ColumnType[SqlInterval, C]

// Json Columns ---------------------------------------------------------------

type JsonColumn[C IConstraint] = ColumnType[SqlJson, C]
type JsonBinaryColumn[C IConstraint] = ColumnType[SqlJsonBinary, C]

// Network Columns ------------------------------------------------------------

type CidrColumn[C IConstraint] = ColumnType[SqlCidr, C]
type INetColumn[C IConstraint] = ColumnType[SqlINet, C]
type MacAddress6Column[C IConstraint] = ColumnType[SqlMacAddress6, C]
type MacAddress8Column[C IConstraint] = ColumnType[SqlMacAddress8, C]

// UUID Columns ---------------------------------------------------------------

type UuidColumn[C IConstraint] = ColumnType[SqlUuid, C]
