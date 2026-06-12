package schema

import "reflect"

type DbResourceId uint

type ForeignKeyRef struct {
	TableId  DbResourceId
	ColumnId DbResourceId
}

type ColumnMeta struct {
	Id DbResourceId

	GoName string
	DbName string

	GoType reflect.Type
	DbType string

	IsPrimary bool
	IsUnique  bool
	IsNotNull bool

	ForeignKey *ForeignKeyRef
}

type TableMeta struct {
	Id DbResourceId

	GoName string
	DbName string

	Columns []ColumnMeta // A column's id is its index
}
