package server

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sauravkuila/broking_setup/pkg/config"
	"github.com/sauravkuila/broking_setup/pkg/dao"
	"github.com/sauravkuila/broking_setup/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

var srv *http.Server
var ctx context.Context
var databases []*gorm.DB

// starts the server with initializations
func Start() error {
	ctx = context.Background()

	config := config.GetConfig()
	logLevel, err := strconv.Atoi(config.GetString("log.Level"))
	if err != nil {
		log.Fatal("Invalid log config: ", err)
	}
	logger.LoggerInit(zapcore.Level(logLevel))
	databases = make([]*gorm.DB, 0)

	postgresConn, err := dao.PsqlConnect()
	if err != nil {
		logger.Log().Error("Failed to connect psql database", zap.Error(err))
		return err
	}

	databases = append(databases, postgresConn)

	return nil
}

// stops the router running in the go routine.
func ShutdownRouter() {
	timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	logger.Log().Info("Shutting down router START")
	defer logger.Log().Info("Shutting down router END")
	if err := srv.Shutdown(timeoutCtx); err != nil {
		logger.Log().Fatal("Server forced to shutdown", zap.Error(err))
	}
	// catching ctx.Done(). timeout of 2 seconds.
	<-timeoutCtx.Done()
	log.Println("timeout of 2 seconds.")
}

// closes all database connections
//
//	closes each database connection made which are saved globally
//	logs error if unable to close
//		function used: *sql.DB.Close()
func CloseDatabase() {
	logger.Log().Info("disconnecting databases START")
	defer logger.Log().Info("disconnecting databases END")
	for _, database := range databases {
		db, _ := database.DB()
		if db != nil {
			err := db.Close()
			if err != nil {
				logger.Log().Error("unable to close db", zap.Error(err))
			}
		} else {
			logger.Log().Error("unable to close db as connection is nil")

		}
	}
}
