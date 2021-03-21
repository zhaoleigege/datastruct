package v1

type LogWriter interface {
	Print(data string) error
	Close()
}

type Logger struct {
	Writers []LogWriter
}

func (log *Logger) RegisterWriter(writer LogWriter) {
	log.Writers = append(log.Writers, writer)
}

func (log *Logger) Print(data string) error {
	for _, writer := range log.Writers {
		if err := writer.Print(data); err != nil {
			return err
		}
	}

	return nil
}

func (log *Logger) Close() {
	for _, writer := range log.Writers {
		writer.Close()
	}
}

func NewLogger() *Logger {
	return &Logger{}
}
