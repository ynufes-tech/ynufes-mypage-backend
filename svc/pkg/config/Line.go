package config

import "ynufes-mypage-backend/pkg/setting"

type line struct {
	ClientID     string
	ClientSecret string
	CallbackURI  string
	CipherKey    string
}

var Line line

func init() {
	config := setting.Get()

	Line = line{
		ClientID:     config.ThirdParty.LineLogin.ClientID,
		ClientSecret: config.ThirdParty.LineLogin.ClientSecret,
		CallbackURI:  config.ThirdParty.LineLogin.CallbackURI,
		CipherKey:    config.ThirdParty.LineLogin.CipherKey,
	}
}
