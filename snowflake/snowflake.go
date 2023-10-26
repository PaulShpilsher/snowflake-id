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

	// public
	DefaultEpoch = int64(1577836800000) // Default epocj unix milleseconds Wed Jan 01 2020 00:00:00 GMT+0000
	TwitterEpoch = int64(1288834974657) // Twitter epoch unix milliseconds Thu Nov 04 2010 01:42:54 GMT+0000

)

var ErrInvalidNodeIDArgument = errors.New("invalid node ID")
var ErrInvalidBaseEpochArgument = errors.New("invalid base epoch")

// SnowflakeGenerator interface
type SnowflakeGenerator interface {
	NextID() int64 // generates next Snowflake ID
}

type snowflakeGenerator struct {
	mu        sync.Mutex
	timestamp int64 // last id timestamp in unix milliseconds
	node      int64 // node precalulated as nodeID << nodeIdOffset
	sequence  int64 // last sequence number for the per the same timestamp
	epoch     int64 // unix milliseconds used offset timestamp portion in the snowflake ID  (timestamp - epoch)
}

// NewGenerator creates a new Snowflake ID generator with default epoch.
// nodeID parameter is the machine or work node identifier range 0 to 1023 inclusive
func NewGenerator(nodeID int) (SnowflakeGenerator, error) {
	return NewGeneratorWithEpoch(nodeID, DefaultEpoch)
}

// NewGeneratorWithEpoch creates a new Snowflake ID generator with timeshift.
// nodeID parameter machine or work node identifier range 0 to 1023 inclusive.
// epoch parameter in unix milliseconds used to offset the current timestamp portion in snowflake id, i.e. current time - epoch.
func NewGeneratorWithEpoch(nodeID int, epoch int64) (SnowflakeGenerator, error) {

	if nodeID < 0 || nodeID > maxNodeID {
		return nil, ErrInvalidNodeIDArgument
	}

	if epoch < 0 || epoch > time.Now().UnixMilli() {
		return nil, ErrInvalidBaseEpochArgument
	}

	return &snowflakeGenerator{
		node:  int64(nodeID) << int64(nodeIdOffset),
		epoch: epoch,
	}, nil
}

// NextID returns the next Snowflake ID
func (s *snowflakeGenerator) NextID() int64 {

	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()

	if now != s.timestamp {
		s.timestamp = now
		s.sequence = 0
	} else {
		if s.sequence = (1 + s.sequence) & maxSequence; s.sequence == 0 {
			// sequence overflow
			// restart with new timestamp
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
			s.timestamp = now
		}
	}

	return ((now - s.epoch) << timestampOffset) | s.node | s.sequence
}
