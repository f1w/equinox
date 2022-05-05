package internal

import (
	"log"
	"time"

	"github.com/Kyagara/equinox/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(retry bool, timeout time.Duration, logLevel api.LogLevel) *zap.SugaredLogger {
	zapConfig := zap.NewProductionConfig()

	equinoxOptions := &api.EquinoxConfig{
		Retry:   retry,
		Timeout: timeout,
	}

	zapConfig.Level = zap.NewAtomicLevelAt(zapcore.Level(logLevel))

	logger, err := zapConfig.Build(zap.Fields(zap.Object("equinox", equinoxOptions)))

	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	return logger.Sugar()
}

func (c *InternalClient) Logger(client string) *zap.SugaredLogger {
	return c.logger.With("client", client)
}
