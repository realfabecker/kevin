package domain

type LogType string

const (
	LogNone   LogType = "none"
	LogTool   LogType = "tool"
	LogScript LogType = "script"
	LogEmbed  LogType = "embed"
)

func ParseLogType(s string) LogType {
	c := map[string]LogType{
		"tool":   LogTool,
		"script": LogScript,
		"embed":  LogEmbed,
	}
	e, ok := c[s]
	if !ok {
		return LogScript
	}
	return e
}
