package helper

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rodericusifo/fasttech-skill-test/mini-app/shared/constant"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockConnectionDatabase(dialect constant.DialectDatabaseType) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
	)
	if err != nil {
		logrus.Fatalf("[sqlmock new] %s", err)
	}
	// defer sqlDB.Close()

	var dialector gorm.Dialector
	switch dialect {
	case constant.POSTGRES:
		dialector = postgres.New(postgres.Config{
			Conn:       sqlDB,
			DriverName: string(constant.POSTGRES),
		})
	}

	// open the database
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Fatalf("[gorm open] %s", err)
	}

	return db, mock
}
