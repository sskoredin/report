package telegram_client

type LogLevel uint

var (
	Debug LogLevel = 0
	Info  LogLevel = 1
	Warn  LogLevel = 2
	Error LogLevel = 3
	Fatal LogLevel = 5
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "Debug"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	default:
		return "Info"
	}
}
