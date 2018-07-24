package common

import (
	"github.com/jinzhu/gorm"
	"github.com/sysco-middleware/commander"
)

var (
	// Commander holds the global commander struct
	Commander *commander.Commander

	// Database holds the global grom database struct
	Database *gorm.DB
)
