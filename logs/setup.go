package logs

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/grpclog"
	"os"
)

// CustomFormatter is a logrus formatter that adds a custom prefix for info messages.
type CustomFormatter struct{}

// Format formats the log entry.
func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {
	levelText := entry.Level.String()

	return []byte("[" + levelText + "] " + entry.Time.Format("2006-01-02 15:04:05") + " " + entry.Message + "\n"), nil
}

// Open all log files
func OpenAllLogFiles() {
	// Create info log
	currentDir, err := os.Getwd()
	logsFile, err := OpenLogFile(currentDir + "/logs/logs.log")
	if err != nil {
		log.Fatal(err)
	}

	// Set the formatter and output for info logs
	log.SetFormatter(&CustomFormatter{})
	log.SetOutput(logsFile)

	// Set the log level for the info logger
	log.SetLevel(log.DebugLevel)

	// Open the file for writing grpc logs
	logFile, err := OpenLogFile(currentDir + "/logs/grpc.log")

	// Create a new logger with the log file as the output
	logger := grpclog.NewLoggerV2(logFile, logFile, logFile)

	// Set the logger as the default logger for gRPC
	grpclog.SetLoggerV2(logger)

}

// Open log file
func OpenLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
