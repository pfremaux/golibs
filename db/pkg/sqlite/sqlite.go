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
	err = createSchemaIfNeeded(dbConnection)
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

type Row interface {
	Id() uint
}

func (c *SqliteClient) Insert(row Row) error {
	// return created row
	return c.dbConnection.Create(row).Error
}

func (c *SqliteClient) Update(row Row) error {
	return c.dbConnection.Model(row).Updates(row).Error
}

func (c *SqliteClient) Delete(row Row) error {
	return c.dbConnection.Delete(row).Error
}

func (c *SqliteClient) Get(row Row) error {
	return c.dbConnection.First(row).Error
}

/*
tags := []Tag{}
	dbConnection.Raw("SELECT * from tags").Scan(&tags)
*/
// TODO PFR just for example
type File struct {
	ID   uint   `gorm:"primaryKey"`
	Key  uint32 `gorm:"uniqueIndex"`
	Tags []Tag  `gorm:"many2many:file_tags;"` // generates file_tags table
}

func (f *File) Id() uint { return f.ID }

// generates table tags
type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

func (t Tag) Id() uint { return t.ID }

func createSchemaIfNeeded(dbConnection *gorm.DB) error {
	needsMigrate := false
	if !dbConnection.Migrator().HasTable(&File{}) || !dbConnection.Migrator().HasTable(&Tag{}) {
		needsMigrate = true
	}
	if !needsMigrate {
		return nil
	}
	// Migre les schémas pour tes modèles (ajoute d'autres structs si besoin)
	err := dbConnection.AutoMigrate(&File{}, &Tag{})
	if err != nil {
		return err
	}
	return nil
}
