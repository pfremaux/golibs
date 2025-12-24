package sqlite

import (
	"os"
	"testing"
)

const DB_NAME = "test.db"

func deleteIfExists(dbName string) {
	// Implementation of deleteIfExists function
	if _, err := os.Stat(dbName); !os.IsNotExist(err) {
		os.Remove(dbName)
	}
}
func TestCreate_table_with_constraints(t *testing.T) {
	deleteIfExists(DB_NAME)
	client, err := Connect(DB_NAME, "", "")
	if err != nil {
		t.Error("	 to connect to SQLite database:", err)
		return
	}
	file := File{
		Key:  12345,
		Tags: []Tag{},
	}
	err = client.Insert(&file)
	if err != nil {
		t.Error(err)
	}
	if file.ID == 0 {
		t.Error("Failed to insert file into the database")
		return
	}
	fileToRead := &File{ID: file.ID}
	err = client.Get(fileToRead)
	if err != nil {
		t.Error(err)
	}
	if fileToRead.Key != file.Key {
		t.Error("Failed to read file from the database")
	}
}
