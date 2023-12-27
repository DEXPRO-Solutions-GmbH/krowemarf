package krowemarf

import (
	"fmt"
	"log/slog"
)

// MigrationLogger is meant to be used with the migration library "golang-migrate".
type MigrationLogger struct {
	slog    *slog.Logger
	verbose bool
}

func NewMigrationLogger(slog *slog.Logger, verbose bool) *MigrationLogger {
	return &MigrationLogger{slog: slog, verbose: verbose}
}

func (m *MigrationLogger) Printf(format string, v ...interface{}) {
	m.slog.Info(fmt.Sprintf(format, v...))
}

func (m *MigrationLogger) Verbose() bool {
	return m.verbose
}
