package mocks

import (
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	gormlogger "gorm.io/gorm/logger"
)

type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) GetLogger() zerolog.Logger {
	args := m.Called()
	return args.Get(0).(zerolog.Logger)
}

func (m *LoggerMock) GetDatabaseLogger() gormlogger.Interface {
	args := m.Called()
	return args.Get(0).(gormlogger.Interface)
}
