package gormadp

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"testing"
	"youtube-clone/pkg/repository/gormadp/dbmodels"
)

func NewTestDBConnection(t *testing.T) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	if db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{}); err != nil {
		t.Fatal(err)
	}

	if err = db.AutoMigrate(
		&dbmodels.User{},
		&dbmodels.Subscribe{},
		&dbmodels.Video{},
		&dbmodels.View{},
		&dbmodels.Like{},
	); err != nil {
		t.Fatal(err)
	}

	return db
}
