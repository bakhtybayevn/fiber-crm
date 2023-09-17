package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Declare a global variable to hold the database connection
var (
	DBConn *gorm.DB
)
