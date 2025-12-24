package db

type DbClient[K any, V any] interface {
	Get(table string, key K) (V, error)
	Set(table string, key K, value V) error
	Add(table string, value V) error
	Remove(table string, key K) error
	ListAll(table string) ([]V, error)
	ListPaged(table string, page int, pageSize int) ([]V, error)
	ListPagedWithFilter(table string, page int, pageSize int, filter func(V) bool) ([]V, error)
}

type Db[K any, V any] interface {
	Connect(uri string, login string, password string) (DbClient[K, V], error)
	Disconnect() error
}
