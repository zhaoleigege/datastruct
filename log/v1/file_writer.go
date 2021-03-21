package v1

import "os"

type FileLogger struct {
	file *os.File
}

func (log *FileLogger) Print(data string) error {
	if log.file == nil {
		return nil
	}

	_, err := log.file.WriteString(data)
	return err
}

func (log *FileLogger) Close() {
	if log.file == nil {
		return
	}

	_ = log.file.Close()
}

func NewFileLogger(filePath string) (LogWriter, error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &FileLogger{file: file}, nil
}
