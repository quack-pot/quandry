package schema

import (
	"log"

	"github.com/quack-pot/quandry/internal/core"
)

type SqlInt struct{}
type SqlSmallInt struct{}
type SqlBigInt struct{}

type SqlFloat struct{}
type SqlDouble struct{}

type SqlMoney struct{}

type SqlBoolean struct{}

type SqlText struct{}

type SqlDate struct{}
type SqlTime struct{}
type SqlDateTime struct{}
type SqlDateTimeZone struct{}
type SqlInterval struct{}

type SqlJson struct{}
type SqlJsonBinary struct{}

type SqlCidr struct{}
type SqlINet struct{}
type SqlMacAddress6 struct{}
type SqlMacAddress8 struct{}

type SqlUuid struct{}

func SqlTypeStringFor[T any](dialect core.ISqlDialect) string {
	var empty T
	switch sql_type := any(empty).(type) {
	case SqlInt:
		return dialect.IntType()
	case SqlSmallInt:
		return dialect.SmallIntType()
	case SqlBigInt:
		return dialect.BigIntType()

	case SqlFloat:
		return dialect.FloatType()
	case SqlDouble:
		return dialect.DoubleType()

	case SqlMoney:
		return dialect.MoneyType()

	case SqlBoolean:
		return dialect.BooleanType()

	case SqlText:
		return dialect.TextType()

	case SqlDate:
		return dialect.DateType()
	case SqlTime:
		return dialect.TimeType()
	case SqlDateTime:
		return dialect.DateTimeType()
	case SqlDateTimeZone:
		return dialect.DateTimeZoneType()
	case SqlInterval:
		return dialect.IntervalType()

	case SqlJson:
		return dialect.JsonType()
	case SqlJsonBinary:
		return dialect.JsonBinaryType()

	case SqlCidr:
		return dialect.CidrType()
	case SqlINet:
		return dialect.INetType()
	case SqlMacAddress6:
		return dialect.MacAddress6Type()
	case SqlMacAddress8:
		return dialect.MacAddress8Type()

	case SqlUuid:
		return dialect.UuidType()

	default:
		log.Panicf("Unknown SQL column type: %v", sql_type)
		return ""
	}
}
