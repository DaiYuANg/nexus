package db

import (
	"github.com/samber/lo"
	"github.com/sony/sonyflake"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"nexus/internal/entity"
	"nexus/internal/model"
	"os"
	"time"
)

var Module = fx.Module("db",
	fx.Provide(newDatabase, snowFlakeGenerator),
	fx.Invoke(databaseLifecycle, registerSnowflakeCallback),
)

func newDatabase(config *model.DatabaseConfig) *gorm.DB {
	connection := lo.Must1(config.GetConnection())
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db := lo.Must1(gorm.Open(connection, &gorm.Config{
		Logger: newLogger,
	}))
	return db
}

func registerSnowflakeCallback(db *gorm.DB, s *sonyflake.Sonyflake, logger2 *zap.Logger) error {
	return db.Callback().Create().Before("gorm:create").Register("snowflake", func(db *gorm.DB) {
		err := fillSnowflake(db, s)
		if err != nil {
			logger2.Error(err.Error())
			return
		}
	})
}

func databaseLifecycle(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.FileResource{},
	)
}
