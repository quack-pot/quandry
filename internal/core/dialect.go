package core

type ISqlDialect interface {
	ArgPlaceholder(arg_number uint) string

	Limit(count uint) string

	IntType() string
	SmallIntType() string
	BigIntType() string
	FloatType() string
	DoubleType() string
	MoneyType() string
	BooleanType() string
	TextType() string
	DateType() string
	TimeType() string
	DateTimeType() string
	DateTimeZoneType() string
	IntervalType() string
	JsonType() string
	JsonBinaryType() string
	CidrType() string
	INetType() string
	MacAddress6Type() string
	MacAddress8Type() string
	UuidType() string
}
