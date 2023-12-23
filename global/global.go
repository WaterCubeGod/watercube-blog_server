package global

import (
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/config"
)

var (
	CONFIG        *config.Config
	DB            *gorm.DB
	LOG           *logrus.Logger
	MySqlLog      logger.Interface
	RDB           *redis.Client
	ES            *elastic.Client
	ARTICLE_INDEX string
)
