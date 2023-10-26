package snowflake_test

import (
	"sync"
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

func TestNewGeneratorWithNegativeNodeId(t *testing.T) {
	if _, err := snowflake.NewGenerator(-1); err != snowflake.ErrInvalidNodeIDArgument {
		t.Error("did not get expected error ErrInvalidNodeIDArgument")
	}
}
func TestNewGeneratorWithNodeIdOverflow(t *testing.T) {
	if _, err := snowflake.NewGenerator(1024); err != snowflake.ErrInvalidNodeIDArgument {
		t.Error("did not get expected error ErrInvalidNodeIDArgument")
	}
}

func TestNewGeneratorWithEpoch(t *testing.T) {
	s, err := snowflake.NewGeneratorWithEpoch(10, snowflake.TwitterEpoch)
	if err != nil {
		t.Errorf("failed to create generator. err: %v", err)
	}
	if s == nil {
		t.Error("generator is nil")
	}
}

func TestNewGeneratorWithNegativeEpoch(t *testing.T) {
	if _, err := snowflake.NewGeneratorWithEpoch(0, -1); err != snowflake.ErrInvalidBaseEpochArgument {
		t.Error("did not get expected error ErrInvalidTimeshiftArgument")
	}
}

func TestNewGeneratorWithFutureEpoch(t *testing.T) {
	futureEpoch := time.Now().UnixMilli() + 1000
	if _, err := snowflake.NewGeneratorWithEpoch(0, futureEpoch); err != snowflake.ErrInvalidBaseEpochArgument {
		t.Error("did not get expected error ErrInvalidTimeshiftArgument")
	}
}

func TestConcurrentGenerationUniqueness(t *testing.T) {

	s, _ := snowflake.NewGenerator(0)

	count := 100000
	ch := make(chan int64, count)

	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			ch <- s.NextID()
		}()
	}
	wg.Wait()
	close(ch)

	m := make(map[int64]interface{})
	for {
		id, ok := <-ch
		if !ok {
			break
		}

		if _, ok := m[id]; ok {
			t.Error("not unique id detected")
			break
		}
		m[id] = nil
	}

}
