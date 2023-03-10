package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/constant"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/migration"
)

var (
	AppConfig   *DefaultConfig
	Environment constant.EnvironmentType
	DB *gorm.DB 
)

func ConfigApps() {
	var (
		environment = flag.String("env", "", "input the environment type")
	)

	flag.Parse()

	switch constant.EnvironmentType(*environment) {
	case constant.DEV:
		viper.SetConfigFile("./environments/dev.application.yaml")
	case constant.STAG:
		viper.SetConfigFile("./environments/stag.application.yaml")
	case constant.PROD:
		viper.SetConfigFile("./environments/prod.application.yaml")
	case constant.TEST:
		viper.SetConfigFile("./environments/test.application.yaml")
	default:
		panic(errors.New("input environment type [ dev | stag | prod | test]"))
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	var conf DefaultConfig
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic(err)
	}

	Environment = constant.EnvironmentType(*environment)
	AppConfig = &conf
}

func ConfigureDatabaseSQL(dialect constant.DialectDatabaseType) {
	var (
		ds Datasource
		db  *gorm.DB
		err error
	)

	switch dialect {
	case constant.POSTGRES:
		ds = AppConfig.Database.Postgres
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require search_path=%s",
		ds.Url,
		ds.Username,
		ds.Password,
		ds.DatabaseName,
		ds.Port,
		ds.Schema)

	cfg := &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ds.Schema,
			SingularTable: false,
		},
	}

	if ds.DebugMode {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}

	switch dialect {
	case constant.POSTGRES:
		db, err = gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			logrus.Panic(err)
		}
	}

	// Auto Migration Models
	db.AutoMigrate(migration.AutoMigrateModelList...)

	sqlDb, err := db.DB()
	if err != nil {
		logrus.Panic(err)
	}

	sqlDb.SetConnMaxIdleTime(ds.ConnectionTimeout)
	sqlDb.SetMaxIdleConns(ds.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(ds.MaxOpenConnection)

	DB = db
}
