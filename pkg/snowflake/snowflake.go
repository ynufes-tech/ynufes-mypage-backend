package snowflake

import (
	"github.com/godruoyi/go-snowflake"
	"strconv"
)

type Snowflake int64

func init() {
	snowflake.SetMachineID(305)
}

func NewSnowflake() Snowflake {
	return Snowflake(snowflake.ID())
}

func (s Snowflake) Int64() int64 {
	return int64(s)
}

func (s Snowflake) String() string {
	return strconv.FormatInt(int64(s), 10)
}
