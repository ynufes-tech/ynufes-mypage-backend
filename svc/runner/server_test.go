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
	req, _ := http.NewRequest("GET", "/api/v1/auth/line/state", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Simulate Callback from Line Login
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/auth/line/callback", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)

	// Check redirection
	assert.Equal(t, "/welcome", w.Header().Get("Location"))
	cookie := w.Header().Get("Set-Cookie")
	cookie = cookie[:strings.Index(cookie, ";")]
	assert.Equal(t, "Authorization=", cookie[:14])

	// Simulate User Info Api
	w = httptest.NewRecorder()
	jwtToken := cookie[14:]
	req, _ = http.NewRequest("GET", "/api/v1/user/info", nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"name_first\":\"\",\"name_last\":\"\",\"type\":1,\"profile_icon_url\":\"https://testUserPicture.com\",\"status\":1}", w.Body.String())

	// Simulate Info Update Api(Insufficient Case for initial registration)
	w = httptest.NewRecorder()
	body := `{"name_first":"太郎","name_last":"横国"}`
	req, _ = http.NewRequest("POST", "/api/v1/user/info/update", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)

	// Simulate Info Update Api(Sufficient Case for initial registration)
	w = httptest.NewRecorder()
	body = `{"name_first":"太郎","name_last":"横国","name_first_kana":"タロウ","name_last_kana":"ヨココク","email":"taro_yokokoku@ynu.jp","student_id":"2164027","gender":1}`
	req, _ = http.NewRequest("POST", "/api/v1/user/info/update", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())

	// Send Insufficient request(initial registration)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/user/info", nil)
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Simulate Revalidation with LineLogin
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/auth/line/callback", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)

}
