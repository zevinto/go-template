package global

import (
	"go-template/internal/config"
	"go-template/internal/crontab"
	"go-template/internal/event"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config          *config.Config
	DB              *gorm.DB
	Logger          *zap.Logger
	EventDispatcher *event.Dispatcher
	Crontab         *crontab.Crontab
)
