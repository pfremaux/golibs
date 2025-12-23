package cache

type Cache[K comparable, V any] interface {
	Get(key K) (V, error)
	Set(key K, value V) error
	Delete(key K) error
	Values() ([]V, error)
}
