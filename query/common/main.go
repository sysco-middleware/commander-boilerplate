package common

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var (
	// Database holds the global database struct
	Database *gorm.DB

	// Router holds the global router struct
	Router *mux.Router
)
