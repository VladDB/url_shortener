package sl

import "log/slog"

// func for adding error message to log as key: value
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
