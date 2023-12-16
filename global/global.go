package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	CONFIG *config.Config
	DB     *gorm.DB
	LOG    *logrus.Logger
)