package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DbConn is the connection handle to the database
var (
	DBConn *gorm.DB // gorm help Golang to interact to database
)
