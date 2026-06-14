package schema

import (
	"reflect"

	"github.com/quack-pot/quandry/internal/core"
)

type SchemaRegistry struct {
	Tables     []TableMeta // Table ids are indices
	GoNameToId map[string]DbResourceId
	dialect    core.ISqlDialect
}

func NewSchemaRegistry(dialect core.ISqlDialect) *SchemaRegistry {
	return &SchemaRegistry{
		Tables:     make([]TableMeta, 0),
		GoNameToId: make(map[string]DbResourceId),
		dialect:    dialect,
	}
}

func (sreg *SchemaRegistry) Register(models ...any) {
	next_id := len(sreg.Tables)

	for _, model := range models {
		model_type := reflect.TypeOf(model)
		model_value := reflect.ValueOf(model)

		if model_type.Kind() == reflect.Ptr {
			model_type = model_type.Elem()
			model_value = model_value.Elem()
		}

		meta := TableMeta{
			Id:      DbResourceId(next_id),
			GoName:  model_type.Name(),
			DbName:  "",
			Columns: make([]ColumnMeta, 0),
		}

		sreg.GoNameToId[meta.GoName] = meta.Id
		next_id += 1

		var next_column_id DbResourceId = 0
		for idx := range model_type.NumField() {
			field := model_type.Field(idx)
			field_value := model_value.Field(idx)

			column, is_column := field_value.Interface().(IColumnType)
			if !is_column {
				continue
			}

			column_meta := ColumnMeta{
				Id:     next_column_id,
				GoName: field.Name,
				DbName: "",
				GoType: field.Type,
				DbType: column.SqlType(sreg.dialect),
			}

			column.ApplyConstraints(&column_meta)

			meta.Columns = append(meta.Columns, column_meta)

			next_column_id += 1
		}

		sreg.Tables = append(sreg.Tables, meta)
	}
}
