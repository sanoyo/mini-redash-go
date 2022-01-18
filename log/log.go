package log

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
	Logger, _ = zap.NewProduction()
	// _ = zap.NewProductionConfig()
}

// ref: https://tomokazu-kozuma.com/minimum-setting-method-of-golangs-logger-zap/
func NewProductionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
