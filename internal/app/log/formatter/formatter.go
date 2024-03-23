package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// JSONFormatter formats logs into parsable json
type JSONFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	// The format to use is the same as for time.Format or time.Parse from the standard
	// library.
	// The standard Library already provides a set of predefined format.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}

// LogDataFields uses for set order for fields in logs
type LogDataFields struct {
	Time  string         `json:"time"`
	Level string         `json:"level"`
	Msg   LogDataMessage `json:"msg"`
}

type LogDataMessage struct {
	Message string `json:"message,omitempty"`
	Extras  any    `json:"extras,omitempty"`
	Error   string `json:"error,omitempty"`
	Func    string `json:"func,omitempty"`
	File    string `json:"file,omitempty"`
}

// Format renders a single log entry
func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := LogDataFields{}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339
	}

	if !f.DisableTimestamp {
		data.Time = entry.Time.Format(timestampFormat)
	}
	data.Msg.Message = entry.Message

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	if f.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}

	return b.Bytes(), nil
}
