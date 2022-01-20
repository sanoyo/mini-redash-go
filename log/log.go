package log

import "go.uber.org/zap"

var (
	Logger *zap.Logger
	err    error
)

func InitLogger() error {
	Logger, err = zap.NewProduction()
	if err != nil {
		return err
	}
	// _ = zap.NewProductionConfig()

	return nil
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
