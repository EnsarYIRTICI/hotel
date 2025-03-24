package util

import (
	"log"
	"os"
)

// Logger struct'ı her katmana özgü loglama yapabilmek için yapılandırıldı
type Logger struct {
	layer  string
	logger *log.Logger
}

// NewLogger fonksiyonu, her katmana özel bir logger oluşturur
func NewLogger(layer string) *Logger {
	return &Logger{
		layer:  layer,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Log fonksiyonu belirli bir mesajı ilgili katman adıyla yazar
func (l *Logger) Log(v ...any) {
	l.logger.Println("["+l.layer+"] ", v)
}
