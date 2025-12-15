package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	var err error
	// Use NewProduction for JSON format, NewDevelopment for console friendly
	Log, err = zap.NewProduction() 
	if err != nil {
		panic(err)
	}
	defer Log.Sync()
}