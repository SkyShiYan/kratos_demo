package data

import (
	"helloworld2/internal/conf"
	"time"

	// "helloworld2/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// NewData .
// Data .
type Data struct {
	// TODO warpped database client
	// db *ent.Client
	db *gorm.DB
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper("data", logger)

	dsn := "root:123!@#qwe@tcp(127.0.0.1:3306)/kratos_demo?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// log.Errorf("-----: %v", conf)
	// client, err := ent.Open(
	// 	conf.Database.Driver,
	// 	conf.Database.Source,
	// )

	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}
	// // Run the auto migration tool.
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Errorf("failed creating schema resources: %v", err)
	// 	return nil, nil, err
	// }
	sqlDB, err := client.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	d := &Data{
		db: client,
	}
	return d, func() {
		logger.Print("message", "closing the data resources")
		sqlDB, err := d.db.DB()
		if err != nil {
			logger.Print("message", "DB也报错了", err)
			return
		}
		sqlDB.Close()
	}, nil
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
