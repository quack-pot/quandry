package schema

import (
	"github.com/quack-pot/quandry/internal/core"
)

type IColumnType interface {
	SqlType(dialect core.ISqlDialect) string
	ApplyConstraints(meta *ColumnMeta)
}

type ColumnType[Owner any, T any, C IConstraint] struct {
	IColumnType
}

func (ColumnType[Owner, T, C]) ApplyConstraints(meta *ColumnMeta) {
	var constraint C
	constraint.Constrain(meta)
}

func (ColumnType[Owner, T, C]) SqlType(dialect core.ISqlDialect) string {
	return SqlTypeStringFor[T](dialect)
}

// Integer Columns ------------------------------------------------------------

type IntColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlInt, C]
type SmallIntColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlSmallInt, C]
type BigIntColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlBigInt, C]

// Decimal Columns ------------------------------------------------------------

type FloatColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlFloat, C]
type DoubleColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlDouble, C]

// Monetary Columns -----------------------------------------------------------

type MoneyColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlMoney, C]

// Boolean Columns ------------------------------------------------------------

type BooleanColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlBoolean, C]

// Character Columns ----------------------------------------------------------

type TextColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlText, C]

// Time Columns ---------------------------------------------------------------

type DateColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlDate, C]
type TimeColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlTime, C]
type DateTimeColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlDateTime, C]
type DateTimeZoneColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlDateTimeZone, C]
type IntervalColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlInterval, C]

// Json Columns ---------------------------------------------------------------

type JsonColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlJson, C]
type JsonBinaryColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlJsonBinary, C]

// Network Columns ------------------------------------------------------------

type CidrColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlCidr, C]
type INetColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlINet, C]
type MacAddress6Column[Owner any, C IConstraint] = ColumnType[Owner, SqlMacAddress6, C]
type MacAddress8Column[Owner any, C IConstraint] = ColumnType[Owner, SqlMacAddress8, C]

// UUID Columns ---------------------------------------------------------------

type UuidColumn[Owner any, C IConstraint] = ColumnType[Owner, SqlUuid, C]
