package internal

import (
	"encoding/json"
	"os"
)

type Log struct {
	ID        string `json:"id"`
	Time      string `json:"time"`
	RemoteIP  string `json:"remote_ip"`
	Host      string `json:"host"`
	Method    string `json:"method"`
	URI       string `json:"uri"`
	UserAgent string `json:"user_agent"`
	Status    int    `json:"status"`
	Latency   int    `json:"latency"`
}

type Logs []Log

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) UpsertLogToFile(log Log) error {
	bytes, err := json.Marshal(log)
	if err != nil {
		return err
	}

	file, err := os.OpenFile("/var/log/datapuppy.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) GetLogsFromFile() (Logs, error) {
	file, err := os.Open("/var/log/datapuppy.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs Logs
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&logs); err != nil {
		return nil, err
	}

	return logs, nil
}
