package config

import "ynufes-mypage-backend/pkg/setting"

var JWT jwt

func init() {
	config := setting.Get()
	JWT = jwt{JWTSecret: config.Application.Authentication.JwtSecret}
}

type jwt struct {
	JWTSecret string
}
