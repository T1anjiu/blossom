package db

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化 MySQL 5.6 连接
func InitDB(user, pass, host, port, dbname string) error {
    // MySQL 5.6 建议显式指定 charset 和 parseTime
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, dbname)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    return nil
}
