package db

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
	initErr  error
)

// InitDB 使用 SQLite 初始化全局的 GORM 数据库连接。
func InitDB(dsn string) (*gorm.DB, error) {
	once.Do(func() {
		instance, initErr = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	})
	return instance, initErr
}

// GetDB 获取全局的 GORM 数据库连接实例。
func GetDB() *gorm.DB {
	return instance
}
