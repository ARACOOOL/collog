package logs

const (
	Debug   = 100
	Info    = 200
	Warning = 300
	Error   = 400
	Fatal   = 500
)

const TimeFormat = "2006-01-02 15:04:05"

type Record struct {
	ID        string                 `json:"id"`
	Source    string                 `json:"source"`
	Category  string                 `json:"category"`
	Level     int                    `json:"level"`
	Message   string                 `json:"message"`
	Trace     string                 `json:"trace"`
	Payload   map[string]interface{} `json:"payload"`
	CreatedAt string                 `json:"created_at"`
}
