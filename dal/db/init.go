package db

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB           *gorm.DB
	dblog        ormlog.Interface
	Redis_client *redis.Client
)

func init() {
	dblog = ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		ormlog.Config{
			SlowThreshold: 100 * time.Millisecond, // slow Sql threshold
			LogLevel:      ormlog.Silent,          // silent level
			Colorful:      true,                   // enable colorful print
		},
	)
	initCache()
	initDB()
}

func initCache() {
	fmt.Println(os.Getenv("REDIS_SECRET"))
	fmt.Println("1212")
	cacheConfig := config.CreateConfig("Cache")
	Redis_client = redis.NewClient(&redis.Options{
		Username: cacheConfig.GetString("redis.Username"),
		Addr:     fmt.Sprintf("%s:%s", cacheConfig.GetString("redis.Host"), cacheConfig.GetString("redis.Port")),
		Password: os.Getenv("REDIS_SECRET"),
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true, // 开发环境跳过证书验证（生产建议设置为 false）
		},
		DB: 0,
	})
	if err := Redis_client.Ping(context.Background()).Err(); err != nil {
		klog.Fatalf("init cache failed: %s", err)
	} else {
		klog.Infof("Redis cache initialized successfully at %s:%s", cacheConfig.GetString("redis.Host"), cacheConfig.GetString("redis.Port"))
	}

}

func initDB() {
	rdbmsConfig := config.CreateConfig("DataBase")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		rdbmsConfig.GetString("Mysql.User"),
		rdbmsConfig.GetString("Mysql.Password"),
		rdbmsConfig.GetString("Mysql.Host"),
		rdbmsConfig.GetInt("Mysql.Port"),
		rdbmsConfig.GetString("Mysql.DBName"),
		rdbmsConfig.GetString("Mysql.CharSet"),
		rdbmsConfig.GetBool("Mysql.ParseTime"),
		rdbmsConfig.GetString("Mysql.loc"),
	)
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 dblog,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		klog.Fatalf("init DB err: %v using %s", err, dsn)
	}

	// gorm open telemetry records database queries and reports DBStats metrics.
	if err = DB.Use(otelgorm.NewPlugin()); err != nil {
		klog.Fatal(err)
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		klog.Fatal(err.Error())
	}
	if err := DB.AutoMigrate(&Relation{}); err != nil {
		klog.Fatal(err.Error())
	}

	if err := DB.AutoMigrate(&Post{}); err != nil {
		klog.Fatal(err.Error())
	}

	if err := DB.AutoMigrate(&Comment{}); err != nil {
		klog.Fatal(err.Error())
	}
	sqlDB, err := DB.DB()
	if err != nil {
		klog.Fatal(err)
	}

	if sqlDB == nil {
		klog.Fatal("sqlDB is nil")
	}

	if err := sqlDB.Ping(); err != nil {
		klog.Error(err.Error())
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxIdleConns(rdbmsConfig.GetInt("Mysql.MaxIdleConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetRdbms() *gorm.DB {
	return DB
}
