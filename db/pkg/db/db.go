package db

type DbClient[K any, V any] interface {
	Get(key K) (V, error)
	Set(key K, value V) error
	Add(value V) error
	Remove(key K) error
	ListAll() ([]V, error)
	ListPaged(page int, pageSize int) ([]V, error)
	ListPagedWithFilter(page int, pageSize int, filter func(V) bool) ([]V, error)
}

type Db[K any, V any] interface {
	Connect(uri string, login string, password string) (DbClient[K, V], error)
	Disconnect() error
}
