package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

//Model is a tuned version of gorm.Model
type Model struct {
	ID        uint64     `form:"id" gorm:"primary_key"`
	CreatedAt time.Time  `binding:"-" form:"-"`
	UpdatedAt time.Time  `binding:"-" form:"-"`
	DeletedAt *time.Time `binding:"-" form:"-"`
}

var db *gorm.DB

//SetDB establishes connection to database and saves its handler into db *sqlx.DB
func SetDB(connection string) {
	var err error
	fmt.Println(connection)
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

//GetDB returns database handler
func GetBD() *gorm.DB {
	return db
}

//AutoMigrate runs gorm auto migration
func AutoMigrate() {
	db.AutoMigrate(&User{}, &Advert{})
}
