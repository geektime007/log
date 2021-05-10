package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Info(1, 2, 3)
	Warn(4, 5, 6)
	Error("error msg")
	P("GK").Error("error message")
}
