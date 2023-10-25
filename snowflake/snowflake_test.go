package snowflake_test

import (
	"testing"
	"time"

	"github.com/PaulShpilsher/snowflake-id/snowflake"
)

func TestNewGenerator(t *testing.T) {
	s, err := snowflake.NewGenerator(10)
	if err != nil {
		t.Errorf("failed to create generator. err: %v", err)
	}
	if s == nil {
		t.Error("generator is nil")
	}
}

func TestNewGeneratorNegativeNodeId(t *testing.T) {
	if _, err := snowflake.NewGenerator(-1); err != snowflake.ErrInvalidNodeIDArgument {
		t.Error("did not get expected error ErrInvalidNodeIDArgument")
	}
}
func TestNewGeneratorToobigNodeId(t *testing.T) {
	if _, err := snowflake.NewGenerator(1024); err != snowflake.ErrInvalidNodeIDArgument {
		t.Error("did not get expected error ErrInvalidNodeIDArgument")
	}
}

func TestNewGeneratorWithTimeshift(t *testing.T) {
	s, err := snowflake.NewGeneratorWithTimeshift(10, 10000)
	if err != nil {
		t.Errorf("failed to create generator. err: %v", err)
	}
	if s == nil {
		t.Error("generator is nil")
	}
}

func TestNewGeneratorWithTimeshiftNegativeTimeshift(t *testing.T) {
	if _, err := snowflake.NewGeneratorWithTimeshift(0, -1); err != snowflake.ErrInvalidTimeshiftArgument {
		t.Error("did not get expected error ErrInvalidTimeshiftArgument")
	}
}

func TestNewGeneratorWithTimeshiftFutureTimeshift(t *testing.T) {
	if _, err := snowflake.NewGeneratorWithTimeshift(0, time.Now().UnixMilli()+1000); err != snowflake.ErrInvalidTimeshiftArgument {
		t.Error("did not get expected error ErrInvalidTimeshiftArgument")
	}
}
