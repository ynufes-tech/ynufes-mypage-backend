package operation

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// UserCreateInput fields can be empty. Default Value will be set in pkg/svc/uc/line/auth_line.go
type UserCreateInput struct {
	AccessToken   string
	RefreshToken  string
	LineServiceID string
	DisplayName   string
	PictureURL    string
	StatusMessage string
}

type UserCreateOutput struct {
	NewUser bool
	JWT     string
}

func UserCreate(t *testing.T, r *gin.Engine, input UserCreateInput) UserCreateOutput {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/auth/line/callback", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)
	cookie := w.Header().Get("Set-Cookie")
	cookie = cookie[:strings.Index(cookie, ";")]
	assert.Equal(t, "Authorization=", cookie[:14])
	var newUser bool
	switch w.Header().Get("Location") {
	case "/welcome":
		newUser = true
	case "/":
		newUser = false
	default:
		t.Fatal("Unexpected Location Header")
	}
	return UserCreateOutput{
		NewUser: newUser,
		JWT:     cookie[14:],
	}
}
