package line

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

var stateCache = make(map[string]int64)

func ReqState(c *gin.Context) {
	newState := IssueNewState()
	c.Writer.WriteString(newState)
	go revokeOldStates()
}

// IssueNewState has to be private method in deployment use
func IssueNewState() string {
	var newState string
	newState = strconv.FormatUint(rand.Uint64(), 10)
	for _, duplicate := stateCache[newState]; duplicate; {
		newState = strconv.FormatUint(rand.Uint64(), 10)
	}
	stateCache[newState] = time.Now().Unix()
	return newState
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
