package schema

import (
	"reflect"
	"strings"
	"unicode"

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

func ToSnakeCase(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return value
	}

	var str_builder strings.Builder
	str_builder.Grow(2 * len(value))

	str_builder.WriteRune(unicode.ToLower(rune(value[0])))
	for idx := 1; idx < len(value); idx++ {
		current := rune(value[idx])

		if !unicode.IsUpper(current) {
			str_builder.WriteRune(unicode.ToLower(current))
			continue
		}

		previous := rune(value[idx-1])

		if unicode.IsLower(previous) {
			str_builder.WriteByte('_')
		} else if unicode.IsUpper(previous) {
			next_idx := idx + 1

			if next_idx < len(value) && unicode.IsLower(rune(value[next_idx])) {
				str_builder.WriteByte('_')
			}
		}

		str_builder.WriteRune(unicode.ToLower(current))
	}

	return str_builder.String()
}

func (sreg *SchemaRegistry) Register(models ...any) {
	next_id := len(sreg.Tables)

	for _, model := range models {
		model_type := reflect.TypeOf(model)
		model_value := reflect.ValueOf(model)

		if model_type.Kind() == reflect.Pointer {
			model_type = model_type.Elem()
			model_value = model_value.Elem()
		}

		meta := TableMeta{
			Id:      DbResourceId(next_id),
			GoName:  model_type.Name(),
			DbName:  ToSnakeCase(model_type.Name()),
			Columns: make([]ColumnMeta, 0),
		}

		sreg.GoNameToId[meta.GoName] = meta.Id
		next_id += 1

		var next_column_id DbResourceId = 0
		for idx := range model_type.NumField() {
			field := model_type.Field(idx)
			field_value := model_value.Field(idx)

			if !field_value.CanInterface() { // Avoids errors with private columns
				continue
			}

			column, is_column := field_value.Interface().(IColumnType)
			if !is_column {
				continue
			}

			column_meta := ColumnMeta{
				Id:     next_column_id,
				GoName: field.Name,
				DbName: ToSnakeCase(field.Name),
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
