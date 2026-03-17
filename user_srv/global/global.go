package global

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

type config struct {
	DSN string `json:"dsn"`
}

func loadConfig() (*config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	if cfg.DSN == "" {
		return nil, logger.ErrInvalidTransaction
	}

	return &cfg, nil
}

func init() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("load config.json error: %v", err)
	}

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := cfg.DSN

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

	// Globally mode
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
