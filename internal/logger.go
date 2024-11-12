package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Detail struct {
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

type Log struct {
	Details []Detail `json:"details"`
}

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) UpsertLogToFile(ld *Detail) error {
	// 既存のログを読み込む
	file, err := os.Open("/var/log/datapuppy.json")
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	var log Log
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&log); err != nil && err.Error() != "EOF" {
			return err
		}
	}

	// 新しいログを追加
	log.Details = append(log.Details, *ld)

	// ファイルを再度開いて書き込む
	file, err = os.OpenFile("/var/log/datapuppy.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(log); err != nil {
		return err
	}

	return nil
}

func (l *Logger) GetLogsFromFile() (*Log, error) {
	file, err := os.Open("/var/log/datapuppy.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var log Log
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&log); err != nil {
		return nil, err
	}

	if len(log.Details) == 0 {
		return nil, fmt.Errorf("no logs found")
	}

	return &log, nil
}
