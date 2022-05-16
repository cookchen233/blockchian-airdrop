package dao

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
    "os"
    "strings"
    "time"
)

var connections = make(map[string]*gorm.DB)

type DataBase struct {
    DbName string
}

func (base *DataBase) SetDbName(DbName string) *DataBase {
    base.DbName = DbName
    return base
}

func (base *DataBase) Db() *gorm.DB {
    _, ok := connections[base.DbName]
    if ok {
        return connections[base.DbName]
    }
    dsns := strings.Split(os.Getenv("mysql_dsn"), ";")
    for _, dsn := range dsns{
        info := strings.Split(dsn, ",")
        if base.DbName == ""{
            base.SetDbName(info[0])
        }
        if info[0] == base.DbName{
            dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", info[2], info[3], info[1], info[0])
            db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
                NamingStrategy: schema.NamingStrategy{
                    TablePrefix:   "",
                    SingularTable: true,
                },
                Logger: logger.Default.LogMode(logger.Error),
            })
            if err != nil {
                panic(err)
            }
            sqlDB, err := db.DB()

            // SetMaxIdleConns 设置空闲连接池中连接的最大数量
            sqlDB.SetMaxIdleConns(10)

            // SetMaxOpenConns 设置打开数据库连接的最大数量。
            sqlDB.SetMaxOpenConns(100)

            // SetConnMaxLifetime 设置了连接可复用的最大时间。
            sqlDB.SetConnMaxLifetime(time.Hour)
            connections[base.DbName] = db
            return db
        }

    }
    return &gorm.DB{}
}