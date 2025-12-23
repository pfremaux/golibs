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

/*
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


			c.Set(record[0], record)
		}
		return CsvClient{
			c:    c,
			path: uri,
		}, nil
	}
*/
func Connect(uri string, login string, password string) (CsvClient, error) {
	c := cache.NewSimpleCache[[]string]()
	cc := CsvClient{
		c:    c,
		path: uri,
	}
	cc.refresh()
	return cc, nil
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

func Disconnect() error {
	return nil
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

func (c *CsvClient) ListAll(table string) ([][]string, error) {
	return c.c.Values()
}
func (c *CsvClient) ListPaged(table string, page int, pageSize int) ([][]string, error) {
	all, err := c.c.Values()
	if err != nil {
		return nil, err
	}
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > len(all) {
		end = len(all)
	}
	return all[start:end], nil
}

func (c *CsvClient) ListPagedWithFilter(table string, page int, pageSize int, filter func([]string) bool) ([][]string, error) {
	return c.c.Values()
}
func (c *CsvClient) Sync() error {
	return c.refresh()
}
