package snowflake

import (
	"testing"
)

const TestSize = 1000000

func TestSnowflake(t *testing.T) {
	var d = make(map[int64]bool, TestSize)
	for i := 0; i < TestSize; i++ {
		s := NewSnowflake()
		_, has := d[s.Int64()]
		if has {
			t.Errorf("DUPLICATE DETECTED: %v", s)
		}
		d[s.Int64()] = true
	}
}
