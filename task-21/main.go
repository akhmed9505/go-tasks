package main

import (
	"fmt"
)

/*
Реализовать паттерн проектирования «Адаптер» на любом примере.
Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого, который ожидает клиент.
Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура) и другой,
несовместимый по интерфейсу потребитель — напишите адаптер, который реализует
нужный интерфейс и делегирует вызовы к встроенному объекту.
Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.
*/

// Target
type Logger interface {
	Debug(message string)
	Info(message string)
	Error(message string)
	Warn(message string)
}

// Adaptee
type LegacyLogger struct{}

func (ll *LegacyLogger) Log(message string) {
	fmt.Println("[OLD]", message)
}

// Adapter
type LoggerAdapter struct {
	legacyLogger *LegacyLogger
}

func NewLoggerAdapter(legacyLogger *LegacyLogger) Logger {
	return &LoggerAdapter{legacyLogger}
}

func (la *LoggerAdapter) Debug(message string) {
	la.legacyLogger.Log("[DEBUG] " + message)
}

func (la *LoggerAdapter) Info(message string) {
	la.legacyLogger.Log("[INFO] " + message)
}

func (la *LoggerAdapter) Error(message string) {
	la.legacyLogger.Log("[ERROR] " + message)
}

func (la *LoggerAdapter) Warn(message string) {
	la.legacyLogger.Log("[WARN] " + message)
}

type Client struct {
	Logger Logger
}

func (c *Client) ProcessRequest(requestID string) {
	c.Logger.Debug("Request " + requestID + " started")
	c.Logger.Info("Processing request " + requestID)
	c.Logger.Warn("Slow request detected: " + requestID)
	c.Logger.Error("Request " + requestID + " failed")
}

func main() {
	legacyLogger := &LegacyLogger{}
	loggerAdapter := NewLoggerAdapter(legacyLogger)

	client := &Client{Logger: loggerAdapter}
	client.ProcessRequest("12345")
}
