package wlog

import (
	"fmt"
	"time"
)

type TextFormatter struct {
	IgnoreBasicFields bool
	DisableColors     bool
}

func (f *TextFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		if !f.DisableColors {
			e.Buffer.WriteString(fmt.Sprintf("%s %s%s%s ", e.Time.Format(time.RFC3339), ColorMapping[e.Level], LevelNameMapping[e.Level], Reset))
		} else {
			e.Buffer.WriteString(fmt.Sprintf("%s %s ", e.Time.Format(time.RFC3339), LevelNameMapping[e.Level]))
		}
		if e.File != "" {
			short := e.File
			for i := len(e.File) - 1; i > 0; i-- {
				if e.File[i] == '/' {
					short = e.File[i+1:]
					break
				}
			}
			e.Buffer.WriteString(fmt.Sprintf(" %s:%d ", short, e.Line))
		}
		e.Buffer.WriteString(" ")
	}
	switch e.Format {
	case FmtEmptySeparate:
		e.Buffer.WriteString(fmt.Sprint(e.Args...))
	default:
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}
	e.Buffer.WriteString("\n")
	return nil
}
