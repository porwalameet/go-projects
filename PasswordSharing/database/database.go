package database

import (
	"context"
	"fmt"

	"github.com/porwalameet/go-projects/PasswordSharing/config"
	"github.com/porwalameet/go-projects/PasswordSharing/logger"
	
	
	"gorm.io/driver/postgres" 
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type DbFactory interface {
	InitDB(context.Context) (*gorm.DB, func(), error)
}

type dbFactory struct {
	c             *config.Config
	loggerFactory logger.LoggerFactory
}

func NewFactory(conf *config.Config, loggerFactory logger.LoggerFactory) DbFactory {
	return &dbFactory{
		c:             conf,
		loggerFactory: loggerFactory,
	}
}

func (f *dbFactory) InitDB(c context.Context) (*gorm.DB, func(), error) {
	appLogger, loggerClose, err := f.loggerFactory.NewLogger()
	if err != nil {
		return nil, nil, err
	}

	defer loggerClose()
	conn, err := f.createConnection(appLogger)
	if err != nil {
		appLogger.Error("cannot create db connection", 
			zap.Error(err), 
			zap.String("provider", f.c.Database.Provider),
			)

		return nil, nil, err
	}

	logger := zapgorm2.New(appLogger)
	db, err := gorm.Open(*conn, &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		appLogger.Error("cannot open gorm",
			zap.Error(err),
			zap.String("provider", f.c.Database.Provider),
		)

		return nil, nil, err
	}

	sql, err := db.DB()
	if err != nil {
		appLogger.Error("failed on db get",
			zap.Error(err),
			zap.String("provider", f.c.Database.Provider),
		)

		return nil, nil, err
	}

	dbClose := func() {
		sql.Close()
	}
	return db.WithContext(c), dbClose, nil

}

func (f *dbFactory) createConnection(appLogger *zap.Logger) (*gorm.Dialector, error) {
	appLogger.Debug("creating DB connection",
		zap.String("provider", f.c.Database.Provider),
		)
	switch f.c.Database.Provider {
	case "pg":
		conn := postgres.New(postgres.Config{
			DSN: f.c.Database.ConnectionString,
		})
		return &conn, nil

	case "sqlite":
		conn := sqlite.Open(f.c.Database.ConnectionString)
		return &conn, nil
	default:
		return nil, fmt.Errorf("cannot create %s connection", f.c.Database.Provider)
	}
}
