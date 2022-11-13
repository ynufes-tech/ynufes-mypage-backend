package line

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

var stateCache map[string]int64

func ReqState(c *gin.Context) {
	var newState string
	newState = strconv.FormatUint(rand.Uint64(), 10)
	for _, duplicate := stateCache[newState]; duplicate; {
		newState = strconv.FormatUint(rand.Uint64(), 10)
	}
	c.Writer.WriteString(newState)
	stateCache[newState] = time.Now().Unix()
	go revokeOldStates()
}

func VerifyState(entry string) bool {
	r, res := stateCache[entry]
	if !res {
		return false
	}
	delete(stateCache, entry)
	if time.Now().Unix()-r > 120000 {
		return false
	}
	return true
}

func revokeOldStates() {
	for s, t := range stateCache {
		if time.Now().Unix()-t > 120000 {
			delete(stateCache, s)
		}
	}
}
