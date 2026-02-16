package domain

type LogType string

const (
	LogNone   LogType = "none"
	LogTool   LogType = "tool"
	LogScript LogType = "script"
)

func ParseLogType(s string) LogType {
	c := map[string]LogType{
		"tool":   LogTool,
		"script": LogScript,
	}
	e, ok := c[s]
	if !ok {
		return LogScript
	}
	return e
}
