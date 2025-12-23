package csv

import "github.com/pfremaux/golibs/cache"

type CsvClient struct {
	c cache.Cache
}

func Connect(uri string, login string, password string) (CsvClient, error) {
	return CsvClient{}, nil
}
func Disconnect() error {
	return nil
}
