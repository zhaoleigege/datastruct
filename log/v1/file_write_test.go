package v1

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewLogger()
	defer logger.Close()

	if fileLogger, err := NewFileLogger("file_log.txt"); err != nil {
		panic(err)
	} else {
		logger.RegisterWriter(fileLogger)
	}

	logger.Print("文件记录日志")
}
