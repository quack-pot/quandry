package postgres

import (
	"strconv"

	"github.com/quack-pot/quandry/internal/core"
)

type PostgresDialect struct{}

func NewDialect() core.ISqlDialect {
	return PostgresDialect{}
}

func (PostgresDialect) ArgPlaceholder(arg_number uint) string {
	number_str := strconv.FormatUint(uint64(arg_number), 10)
	return "$" + number_str
}

func (PostgresDialect) Limit(count uint) string {
	number_str := strconv.FormatUint(uint64(count), 10)
	return "LIMIT " + number_str
}

func (PostgresDialect) IntType() string          { return "INTEGER" }
func (PostgresDialect) SmallIntType() string     { return "SMALLINT" }
func (PostgresDialect) BigIntType() string       { return "BIGINT" }
func (PostgresDialect) FloatType() string        { return "REAL" }
func (PostgresDialect) DoubleType() string       { return "DOUBLE PRECISION" }
func (PostgresDialect) MoneyType() string        { return "MONEY" }
func (PostgresDialect) BooleanType() string      { return "BOOLEAN" }
func (PostgresDialect) TextType() string         { return "TEXT" }
func (PostgresDialect) DateType() string         { return "DATE" }
func (PostgresDialect) TimeType() string         { return "TIME" }
func (PostgresDialect) DateTimeType() string     { return "TIMESTAMP" }
func (PostgresDialect) DateTimeZoneType() string { return "TIMESTAMP WITH TIME ZONE" }
func (PostgresDialect) IntervalType() string     { return "INTERVAL" }
func (PostgresDialect) JsonType() string         { return "JSON" }
func (PostgresDialect) JsonBinaryType() string   { return "JSONB" }
func (PostgresDialect) CidrType() string         { return "CIDR" }
func (PostgresDialect) INetType() string         { return "INET" }
func (PostgresDialect) MacAddress6Type() string  { return "MACADDR" }
func (PostgresDialect) MacAddress8Type() string  { return "MACADDR8" }
func (PostgresDialect) UuidType() string         { return "UUID" }
