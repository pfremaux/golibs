package csv

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/pfremaux/golibs/cache/pkg/cache"
)

type CsvClient struct {
	path string
	c    cache.Cache[string, []string]
}

var EMPTY = CsvClient{}

func Load(uri string) (CsvClient, error) {
	c := cache.NewSimpleCache[[]string]()
	cc := CsvClient{
		c:    c,
		path: uri,
	}
	err := cc.Sync()
	return cc, err
}

func (c *CsvClient) Get(table string, key string) ([]string, error) {
	return c.c.Get(key)
}

func (c *CsvClient) Set(table string, key string, value []string) error {
	return fmt.Errorf("Unsupported operation Set with CSV client. It's a readonly lib.")
}

func (c *CsvClient) Remove(table string, key string) error {
	return c.c.Delete(key)
}

func (c *CsvClient) ListAll() ([][]string, error) {
	return c.c.Values()
}

func (c *CsvClient) Sync() error {
	return c.refresh()
}

func (cc *CsvClient) refresh() error {
	c := cache.NewSimpleCache[[]string]()
	file, err := os.Open(cc.path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		if len(record) < 2 {
			continue // Skip invalid records
		}

		c.Set(record[0], record)
	}
	cc.c = c
	return nil
}
