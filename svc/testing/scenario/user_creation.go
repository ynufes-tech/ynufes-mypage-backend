package scenario

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
	userModel "ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/schema/user"
	"ynufes-mypage-backend/svc/testing/operation"
)

// TestUserCreation simulates following steps
// 1. Create a new user
// 2. Get user info
// 3. Create a new user with same LineServiceID
// 4. Get user info
// 5. Create a new user with different LineServiceID
// 6. Get user info
func TestUserCreation(t *testing.T, r *gin.Engine) {
	// 1. Create a new user
	ipt1 := operation.UserCreateInput{
		AccessToken:   "accessToken1",
		RefreshToken:  "refreshToken1",
		LineServiceID: "TestLineServiceID1",
		DisplayName:   "displayName1",
		PictureURL:    "https://test.com/hehe.jpg",
		StatusMessage: "StatusMessage1",
	}
	out1 := operation.UserCreate(t, r, ipt1)
	assert.Equal(t, true, out1.NewUser)

	// 2. Get user info
	infoOut1 := operation.UserInfo(t, r, operation.UserInfoInput{
		Authorization: out1.JWT,
	})
	if !assert.Equal(t, 200, infoOut1.Code) {
		t.Errorf("TestUserCreation Failed: IncorrectCode (%d ,%d)", 200, infoOut1.Code)
		return
	}
	infoOut1.StrictValidate(t, user.InfoResponse{
		NameFirst:       "",
		NameLast:        "",
		Type:            int(userModel.StatusRegistered),
		ProfileImageURL: ipt1.PictureURL,
		Status:          int(userModel.StatusNew),
	})

	// 3. Create a new user with same LineServiceID
	ipt2 := operation.UserCreateInput{
		AccessToken:   "accessToken2",
		RefreshToken:  "refreshToken2",
		LineServiceID: ipt1.LineServiceID,
		DisplayName:   "displayName2",
		PictureURL:    "https://test.com/hehe2.jpg",
		StatusMessage: "StatusMessage2",
	}
	out2 := operation.UserCreate(t, r, ipt2)
	assert.Equal(t, false, out2.NewUser)

	// 4. Get user info
	infoOut2 := operation.UserInfo(t, r, operation.UserInfoInput{
		Authorization: out2.JWT,
	})
	if !assert.Equal(t, 200, infoOut2.Code) {
		t.Errorf("TestUserCreation Failed: IncorrectCode (%d ,%d)", 200, infoOut2.Code)
		return
	}
	infoOut2.StrictValidate(t, user.InfoResponse{
		NameFirst:       "",
		NameLast:        "",
		Type:            int(userModel.StatusRegistered),
		ProfileImageURL: ipt2.PictureURL,
		Status:          int(userModel.StatusNew),
	})

	// 5. Create a new user with different LineServiceID
	ipt3 := operation.UserCreateInput{
		AccessToken:   "accessToken3",
		RefreshToken:  "refreshToken3",
		LineServiceID: "TestLineServiceID3",
		DisplayName:   "displayName3",
		PictureURL:    "https://test.com/hehe3.jpg",
		StatusMessage: "StatusMessage3",
	}
	out3 := operation.UserCreate(t, r, ipt3)
	assert.Equal(t, true, out3.NewUser)

	// 6. Get user info
	infoOut3 := operation.UserInfo(t, r, operation.UserInfoInput{
		Authorization: out3.JWT,
	})
	if !assert.Equal(t, 200, infoOut3.Code) {
		t.Errorf("TestUserCreation Failed: IncorrectCode (%d ,%d)", 200, infoOut3.Code)
		return
	}
	infoOut3.StrictValidate(t, user.InfoResponse{
		NameFirst:       "",
		NameLast:        "",
		Type:            int(userModel.StatusRegistered),
		ProfileImageURL: ipt3.PictureURL,
		Status:          int(userModel.StatusNew),
	})
}
