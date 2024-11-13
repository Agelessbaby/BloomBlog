package db

import (
	"fmt"
	"github.com/Agelessbaby/BloomBlog/util/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB    *gorm.DB
	dblog ormlog.Interface
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
	initDB()
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
		klog.Error(err.Error())
	}
	if err := DB.AutoMigrate(&Relation{}); err != nil {
		klog.Error(err.Error())
	}
	//TODO ADD other auto migrate
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

	sqlDB.SetMaxIdleConns(rdbmsConfig.GetInt("Mysql.MaxIdleConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetRdbms() *gorm.DB {
	return DB
}
