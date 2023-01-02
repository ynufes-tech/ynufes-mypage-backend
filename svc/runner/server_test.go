package runner

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestImplement(t *testing.T) {
	r := gin.Default()
	w := httptest.NewRecorder()
	rg := r.Group("/api/v1")
	err := Implement(rg)
	assert.NoError(t, err)
	req, _ := http.NewRequest("GET", "/api/v1/line/state", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Simulate Callback from Line Login
	req, _ = http.NewRequest("GET", "/api/v1/auth/line/callback", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)

	// Check redirection
	assert.Equal(t, "/welcome", w.Header().Get("Location"))
	cookie := w.Header().Get("Set-Cookie")
	assert.Equal(t, "Authorization=", cookie[:14])
	jwtToken := cookie[14:]
	req, _ = http.NewRequest("GET", "/api/v1/user/info", nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name_first\":\"\",\"name_last\":\"\",\"type\":1,\"profile_icon_url\":\"https://testUserPicture.com\",\"status\":1}", w.Body.String())

	// Simulate user info update
	body := `{"name_first":"太郎"","name_last":"横国","type":0,"profile_icon_url":"https://testUserPicture.com","status":0}`
	req, _ = http.NewRequest("POST", "/api/v1/user/info/update", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())

	req, _ = http.NewRequest("GET", "/api/v1/user/info", nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name_first\":\"太郎\",\"name_last\":\"横国\",\"type\":1,\"profile_icon_url\":\"https://testUserPicture.com\",\"status\":2}", w.Body.String())

	// Simulate Revalidation with LineLogin
	req, _ = http.NewRequest("GET", "/api/v1/auth/line/callback", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)

}
