package core

type IDatabase interface {
	RegisterTable(models ...any)
	Close() error

	IDbTransactionFactory
}
