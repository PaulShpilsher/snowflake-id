package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	timestampBits = uint8(41) // 41 bit timestamp
	nodeIdBits    = uint8(10) // 10 bit node ID
	sequenceBits  = uint8(12) // 12 bit sequence

	maxTimestamp = (1 << timestampBits) - 1       // 22 value value bit-mask = 0x1ffffffffff
	maxNodeID    = (1 << nodeIdBits) - 1          // 10 bit value bit-mask = 0x3ff
	maxSequence  = int64((1 << sequenceBits) - 1) // 12 value value bit-mask = 0xfff

	nodeIdOffset    = sequenceBits              // Node ID offset to the left
	timestampOffset = nodeIdBits + sequenceBits // The timestamp offset to the left
)

var ErrInvalidNodeIDArgument = errors.New("invalid node ID")
var ErrInvalidTimeshiftArgument = errors.New("invalid timeshift")

// SnowflakeGenerator interface
type SnowflakeGenerator interface {
	NextID() int64 // generates next Snowflake ID
}

type snowflake struct {
	mu        sync.Mutex
	timestamp int64 // last id timestamp in unix milliseconds
	nodeID    int64 // node id
	sequence  int64 // last sequence number for the per the same timestamp

	timeshift int64 // unix milliseconds used in timestamp portion if the snowflake ID as (curent timestamp - timeshift)

}

// NewGenerator creates a new Snowflake ID generator.
// nodeID parameter is the machine or work node identifier range 0 to 1023 inclusive
func NewGenerator(nodeID int) (SnowflakeGenerator, error) {
	return NewGeneratorWithTimeshift(nodeID, 0)
}

// NewGeneratorWithTimeshift creates a new Snowflake ID generator with timeshift.
// nodeID parameter is the machine or work node identifier range 0 to 1023 inclusive.
// timeshift parameter is in unix milliseconds used to offset the current timestamp portion in snowflake id, i.e. current time - timeshif
func NewGeneratorWithTimeshift(nodeID int, timeshift int64) (SnowflakeGenerator, error) {

	if nodeID < 0 || nodeID > maxNodeID {
		return nil, ErrInvalidNodeIDArgument
	}

	if timeshift < 0 || timeshift > getUnixMilli() {
		return nil, ErrInvalidTimeshiftArgument
	}

	return &snowflake{
		nodeID:    int64(nodeID),
		timeshift: timeshift,
	}, nil
}

// NextID returns the next Snowflake ID
func (s *snowflake) NextID() int64 {

	s.mu.Lock()
	defer s.mu.Unlock()

	now := getUnixMilli()

	if now != s.timestamp {
		s.timestamp = now
		s.sequence = 0
	} else {
		if s.sequence = (1 + s.sequence) & maxSequence; s.sequence == 0 {
			// sequence overflow
			for now <= s.timestamp {
				now = getUnixMilli()
			}
			s.timestamp = now
		}
	}

	return ((now - s.timeshift) << timestampOffset) | (s.nodeID << nodeIdOffset) | s.sequence
}

// returns current milliseconds
func getUnixMilli() int64 {
	return time.Now().UnixMilli()
}
