package operation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserInfoUpdateInput struct {
	Authorization string
}

type UserInfoUpdateOutput struct {
	Code        int
	ErrResponse string
}

func UserInfoUpdate(t *testing.T, r *gin.Engine, input UserInfoUpdateInput) UserInfoUpdateOutput {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/api/v1/user", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", input.Authorization))
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		return UserInfoUpdateOutput{
			Code:        w.Code,
			ErrResponse: w.Body.String(),
		}
	}
	return UserInfoUpdateOutput{
		Code: w.Code,
	}
}
