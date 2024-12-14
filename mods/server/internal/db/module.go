package db

import (
	"github.com/samber/lo"
	"github.com/sony/sonyflake"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/config"
	"nexus/internal/entity"
)

var Module = fx.Module("db",
	fx.Provide(newDatabase, snowFlakeGenerator),
	fx.Invoke(databaseLifecycle, registerSnowflakeCallback),
)

type DatabaseParam struct {
	fx.In
	Config *config.DatabaseConfig
	*zap.Logger
}

func newDatabase(param DatabaseParam) *gorm.DB {
	config, logger := param.Config, param.Logger
	connection := lo.Must1(config.GetConnection())
	db := lo.Must1(gorm.Open(connection, &gorm.Config{
		Logger: newGormZapAdapter(logger),
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
		&entity.File{},
		&entity.Folder{},
		&entity.UserGroup{},
	)
}
