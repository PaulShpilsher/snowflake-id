package snowflake_test

import (
	"testing"

	"github.com/PaulShpilsher/snowflake-id/snowflake"
)

func TestNewGenerator(t *testing.T) {
	s, err := snowflake.NewGenerator(10)
	if err != nil {
		t.Error("Failed to create generator", err)
	}
	if s == nil {
		t.Error("Created generator is nil")
	}
}
