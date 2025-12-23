package csv

import (
	"encoding/csv"
	"os"

	"github.com/pfremaux/golibs/cache/pkg/cache"
)

type CsvClient struct {
	c cache.Cache[string, []string]
}

var EMPTY = CsvClient{}

func Connect(uri string, login string, password string) (CsvClient, error) {
	c := cache.NewSimpleCache[[]string]()
	file, err := os.Open(uri)
	if err != nil {
		return EMPTY, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return EMPTY, err
	}

	for _, record := range records {
		if len(record) < 2 {
			continue // Skip invalid records
		}
		//id, err := strconv.ParseUint(record[0], 10, 32)
		/*if err != nil {
			continue // Skip invalid ID
		}*/
		c.Set(record[0], record)
	}
	return CsvClient{
		c: c,
	}, nil
}
func Disconnect() error {
	return nil
}

func (c *CsvClient) Get(key string) ([]string, error) {
	return c.c.Get(key)
}

func (c *CsvClient) Set(key string, value []string) error {
	return c.c.Set(key, value)
}

func (c *CsvClient) Remove(key string) error {
	return c.c.Delete(key)
}

func (c *CsvClient) ListAll() ([][]string, error) {
	return c.c.Values(), nil
}
func (c *CsvClient) ListPaged(page int, pageSize int) ([][]string, error) {
	return c.c.Values(), nil
}
func (c *CsvClient) ListPagedWithFilter(page int, pageSize int, filter func([]string) bool) ([][]string, error) {
	return c.c.Values(), nil
}
func (c *CsvClient) Sync() error {
	return nil
}
