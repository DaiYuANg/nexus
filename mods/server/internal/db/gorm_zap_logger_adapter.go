package db

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"time"
)

type gormZapAdapter struct {
	*zap.Logger
}

func (g *gormZapAdapter) LogMode(level logger.LogLevel) logger.Interface {
	// 这里返回一个新的 ZapGormLogger 以支持链式设置
	return g
}

func (g *gormZapAdapter) Info(ctx context.Context, s string, i ...interface{}) {
	// 使用 Zap 记录 info 日志
	g.Logger.Sugar().Infof(s, i...)
}

func (g *gormZapAdapter) Warn(ctx context.Context, s string, i ...interface{}) {
	g.Logger.Sugar().Warnf(s, i...)
}

func (g *gormZapAdapter) Error(ctx context.Context, s string, i ...interface{}) {
	g.Logger.Sugar().Errorf(s, i...)
}

func (g *gormZapAdapter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	g.Logger.Sugar().Debug("Debug")
}

func newGormZapAdapter(logger *zap.Logger) *gormZapAdapter {
	return &gormZapAdapter{logger}
}
