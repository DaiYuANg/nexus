package model

import (
	"bytes"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
)

const (
	MysqlDSNTemplate      = "{{.Username}}:{{.Password}}@tcp({{.Host}}:{{.Port}})/{{.DBName}}?charset=utf8mb4&parseTime=True&loc=Local"
	PostgresqlDSNTemplate = "host={{.Host}} user={{.Username}} password={{.Password}} dbname={{.DBName}} port={{.Port}} sslmode=disable TimeZone=Asia/Shanghai"
)

type DatabaseConfig struct {
	Type     string `koanf:"type"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	DBName   string `koanf:"dbname"`
	Path     string `koanf:"path"`
}

func (c *DatabaseConfig) GetConnection() (gorm.Dialector, error) {
	var dsn string
	var err error

	switch c.Type {
	case "mysql":
		dsn, err = c.getDSN(MysqlDSNTemplate)
		if err != nil {
			return nil, err
		}
		return mysql.Open(dsn), nil
	case "postgres":
		dsn, err = c.getDSN(PostgresqlDSNTemplate)
		if err != nil {
			return nil, err
		}
		return postgres.Open(dsn), nil
	case "sqlite":
		return sqlite.Open(c.Path), nil
	default:
		// 返回一个错误，指示数据库类型未知
		return nil, errors.New("unknown database type")
	}
}

// 获取数据库的连接字符串
func (c *DatabaseConfig) getDSN(templateStr string) (string, error) {
	t, err := template.New("dsn").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var output bytes.Buffer
	err = t.Execute(&output, c)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
