package debugger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("check for panics", func(t *testing.T) {
		assert.NotPanics(t, func() { NewDebugger().Close() })
	})
}
