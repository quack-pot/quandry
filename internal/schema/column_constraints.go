package schema

type Constraint interface {
	Constrain(*ColumnMeta)
}

// None -----------------------------------------------------------------------

type None struct{}

func (None) Constrain(*ColumnMeta) { /* No-Op */ }

// PrimaryKey -----------------------------------------------------------------

type PrimaryKey struct{}

func (PrimaryKey) Constrain(meta *ColumnMeta) {
	meta.IsPrimary = true
}

// Unique ---------------------------------------------------------------------

type Unique struct{}

func (Unique) Constrain(meta *ColumnMeta) {
	meta.IsUnique = true
}

// Required -------------------------------------------------------------------

type Required struct{}

func (Required) Constrain(meta *ColumnMeta) {
	meta.IsNotNull = true
}

// ForeignKey -----------------------------------------------------------------

type ForeignKey[T any] struct{}

func (ForeignKey[T]) Constrain(meta *ColumnMeta) {}

// AutoIncrement --------------------------------------------------------------

type AutoIncrement struct{}

func (AutoIncrement) Constrain(meta *ColumnMeta) {}

// Computed -------------------------------------------------------------------

type Computed struct{}

func (Computed) Constrain(meta *ColumnMeta) {}

// Constraints ----------------------------------------------------------------

type Constraints[A Constraint, B Constraint] struct{}

func (Constraints[A, B]) Constrain(meta *ColumnMeta) {
	var constraintA A
	var constraintB B

	constraintA.Constrain(meta)
	constraintB.Constrain(meta)
}

// PrimaryKeyAuto -------------------------------------------------------------

type PrimaryKeyAuto Constraints[PrimaryKey, AutoIncrement]
