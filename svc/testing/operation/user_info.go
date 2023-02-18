package operation

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"ynufes-mypage-backend/svc/pkg/schema/user"
)

type UserInfoInput struct {
	Authorization string
}

type UserInfoOutput struct {
	Code        int
	Response    *user.InfoResponse
	ErrResponse string
}

func UserInfo(t *testing.T, r *gin.Engine, input UserInfoInput) UserInfoOutput {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/user", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", input.Authorization))
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		return UserInfoOutput{
			Code:        w.Code,
			Response:    nil,
			ErrResponse: w.Body.String(),
		}
	}
	var info user.InfoResponse
	err := json.Unmarshal(w.Body.Bytes(), &info)
	assert.NoError(t, err)
	return UserInfoOutput{
		Code:     w.Code,
		Response: &info,
	}
}

func (o UserInfoOutput) StrictValidate(t *testing.T, suppose user.InfoResponse) {
	assert.Equal(t, suppose.NameFirst, o.Response.NameFirst)
	assert.Equal(t, suppose.NameLast, o.Response.NameLast)
	assert.Equal(t, suppose.Type, o.Response.Type)
	assert.Equal(t, suppose.ProfileImageURL, o.Response.ProfileImageURL)
	assert.Equal(t, suppose.Status, o.Response.Status)
}
