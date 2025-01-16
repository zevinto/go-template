package global

import (
	"go-template/internal/config"

	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
