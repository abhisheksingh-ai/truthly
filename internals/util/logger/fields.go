package logger

import "log/slog"

func Str(k, v string) slog.Attr {
	return slog.String(k, v)
}

func Int(k string, v int) slog.Attr {
	return slog.Int(k, v)
}

func Err(err error) slog.Attr {
	return slog.Any("error", err)
}

func Any(k string, v any) slog.Attr {
	return slog.Any(k, v)
}
