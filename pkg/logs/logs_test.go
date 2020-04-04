package logs

import (
	"testing"
)

func TestLogs(t *testing.T) {
	Logger.Errorf("test %v", "1231")
	t.Logf("test")
}
