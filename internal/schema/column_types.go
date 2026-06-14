package schema

// Integer Columns ------------------------------------------------------------

type IntColumn[C Constraint] struct{}
type SmallIntColumn[C Constraint] struct{}
type BigIntColumn[C Constraint] struct{}

// Decimal Columns ------------------------------------------------------------

type FloatColumn[C Constraint] struct{}
type DoubleColumn[C Constraint] struct{}

// Monetary Columns -----------------------------------------------------------

type MoneyColumn[C Constraint] struct{}

// Boolean Columns ------------------------------------------------------------

type BooleanColumn[C Constraint] struct{}

// Character Columns ----------------------------------------------------------

type TextColumn[C Constraint] struct{}

// Time Columns ---------------------------------------------------------------

type DateColumn[C Constraint] struct{}
type TimeColumn[C Constraint] struct{}
type DateTimeColumn[C Constraint] struct{}

type DateZoneColumn[C Constraint] struct{}
type TimeZoneColumn[C Constraint] struct{}
type DateTimeZoneColumn[C Constraint] struct{}

type IntervalColumn[C Constraint] struct{}

// Json Columns ---------------------------------------------------------------

type JsonColumn[C Constraint] struct{}
type JsonBinaryColumn[C Constraint] struct{}

// Network Columns ------------------------------------------------------------

type CidrColumn[C Constraint] struct{}
type INetColumn[C Constraint] struct{}
type MacAddress6Column[C Constraint] struct{}
type MacAddress8Column[C Constraint] struct{}

// UUID Columns ---------------------------------------------------------------

type UuidColumn[C Constraint] struct{}
