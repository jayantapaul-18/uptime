package environmentvar

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEnv(t *testing.T) {
	os.Setenv("NODE_ENV", "production")
	defer os.Unsetenv("NODE_ENV")

	actual := os.Getenv("NODE_ENV")
	assert.Equal(t, actual, "production")
}
