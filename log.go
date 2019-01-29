package dexcom

type Logger interface {
	Log(keyvals ...interface{}) error
}

type noOpLogger struct {}

func (l noOpLogger) Log(keyvals ...interface{}) error {
	return nil
}