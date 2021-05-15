package vivliothiki

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEnv(t *testing.T) {
	t.Run("should return as development", func(t *testing.T) {
		envStage := GetEnv()
		assert.Equal(t, EnvDevelopment, envStage)
	})
}
