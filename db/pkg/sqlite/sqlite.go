package sqlite

import (
	"github.com/pfremaux/golibs/cache/pkg/cache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteClient struct {
	dbConnection *gorm.DB
	c            map[string]cache.Cache[string, any]
}

func Connect(uri string, login string, password string) (SqliteClient, error) {
	dbConnection, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		return SqliteClient{}, err
	}
	c := make(map[string]cache.Cache[string, any])
	cc := SqliteClient{
		dbConnection: dbConnection,
		c:            c,
	}
	return cc, err
}

/*
tags := []Tag{}
	dbConnection.Raw("SELECT * from tags").Scan(&tags)
*/
