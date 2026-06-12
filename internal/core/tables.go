package core

import (
	"iter"
	"reflect"
	"sync"
)

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

	Columns    []ColumnMeta
	IdToColumn map[DbResourceId]int
}

type SchemaRegistry struct {
	tables []TableMeta

	idxToId     map[int]DbResourceId
	idToTable   map[DbResourceId]int
	nextTableId DbResourceId

	mutex sync.RWMutex
}

func NewSchemaRegistry() *SchemaRegistry {
	return &SchemaRegistry{
		tables: make([]TableMeta, 0),

		idxToId:   make(map[int]DbResourceId),
		idToTable: make(map[DbResourceId]int),

		nextTableId: 0,
		mutex:       sync.RWMutex{},
	}
}

func (sreg *SchemaRegistry) Add(table TableMeta) DbResourceId {
	sreg.mutex.Lock()
	defer sreg.mutex.Unlock()

	id := sreg.nextTableId
	index := len(sreg.tables)

	sreg.nextTableId += 1
	sreg.tables = append(sreg.tables, table)

	sreg.idxToId[index] = id
	sreg.idToTable[id] = index

	return id
}

func (sreg *SchemaRegistry) Remove(id DbResourceId) {
	sreg.mutex.Lock()
	defer sreg.mutex.Unlock()

	index, has_id := sreg.idToTable[id]
	if !has_id {
		return
	}

	last_index := len(sreg.tables) - 1
	last_id := sreg.idxToId[last_index]

	delete(sreg.idToTable, id)
	delete(sreg.idxToId, last_index)

	if index == last_index {
		sreg.tables = sreg.tables[:last_index]
		return
	}

	sreg.tables[index] = sreg.tables[last_index]
	sreg.tables = sreg.tables[:last_index]

	sreg.idToTable[last_id] = index
	sreg.idxToId[index] = last_id
}

func (sreg *SchemaRegistry) Get(id DbResourceId) (TableMeta, bool) {
	sreg.mutex.RLock()
	defer sreg.mutex.RUnlock()

	index, has_id := sreg.idToTable[id]

	if !has_id {
		return TableMeta{}, false
	}

	return sreg.tables[index], true
}

func (sreg *SchemaRegistry) All() iter.Seq[*TableMeta] {
	return func(yield func(*TableMeta) bool) {
		sreg.mutex.RLock()
		defer sreg.mutex.RUnlock()

		for idx := range sreg.tables {
			if !yield(&sreg.tables[idx]) {
				return
			}
		}
	}
}
