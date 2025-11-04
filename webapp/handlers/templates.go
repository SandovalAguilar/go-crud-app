package handlers

import (
	"html/template"
	"strings"
	"time"
)

var Templates = template.Must(template.New("").Funcs(template.FuncMap{
	"contains": strings.Contains,
	"sub": func(a, b int) int {
		return a - b
	},
	"formatDate": func(date interface{}) string {
		switch v := date.(type) {
		case time.Time:
			// If it's a time.Time, format it as YYYY-MM-DD
			return v.Format("2006-01-02")
		case *time.Time:
			// If it's a pointer to time.Time (nullable date)
			if v != nil {
				return v.Format("2006-01-02")
			}
			return ""
		case string:
			// If it's already a string, try to parse and reformat
			if t, err := time.Parse(time.RFC3339, v); err == nil {
				return t.Format("2006-01-02")
			}
			// Try parsing without timezone
			if t, err := time.Parse("2006-01-02T15:04:05Z07:00", v); err == nil {
				return t.Format("2006-01-02")
			}
			// If it's already in correct format, return as-is
			return v
		default:
			return ""
		}
	},
}).ParseGlob("templates/**/*"))
