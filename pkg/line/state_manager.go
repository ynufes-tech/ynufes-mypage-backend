package line

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

type AuthStateManager struct {
	stateCache map[string]int64
}

func NewAuthStateManager() AuthStateManager {
	return AuthStateManager{
		stateCache: make(map[string]int64),
	}
}

func ReqState(c *gin.Context) {
	m := NewAuthStateManager()
	newState := m.IssueNewState()
	c.Writer.WriteString(newState)
	go m.RevokeOldStates()
}

// IssueNewState has to be private method in deployment use
func (m AuthStateManager) IssueNewState() string {
	var newState string
	newState = strconv.FormatUint(rand.Uint64(), 10)
	for _, duplicate := m.stateCache[newState]; duplicate; {
		newState = strconv.FormatUint(rand.Uint64(), 10)
	}
	m.stateCache[newState] = time.Now().Unix()
	return newState
}

func (m AuthStateManager) VerifyState(entry string) bool {
	r, res := m.stateCache[entry]
	if !res {
		return false
	}
	delete(m.stateCache, entry)
	if time.Now().Unix()-r > 120000 {
		return false
	}
	return true
}

func (m AuthStateManager) RevokeOldStates() {
	for s, t := range m.stateCache {
		if time.Now().Unix()-t > 120000 {
			delete(m.stateCache, s)
		}
	}
}
