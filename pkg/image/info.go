package image

import (
	"bytes"
	"golang.org/x/image/webp"
	"image/jpeg"
	"image/png"
)

type JPEGInfo struct{}
type PNGInfo struct{}
type WEBPInfo struct{}

func NewJPEGInfo() JPEGInfo {
	return JPEGInfo{}
}

func NewPNGInfo() PNGInfo {
	return PNGInfo{}
}

func NewWEBPInfo() WEBPInfo {
	return WEBPInfo{}
}

func (p JPEGInfo) ExtractInfo(data []byte) (width, height int, err error) {
	config, err := jpeg.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return 0, 0, err
	}
	return config.Width, config.Height, nil
}

func (p PNGInfo) ExtractInfo(data []byte) (width, height int, err error) {
	config, err := png.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return 0, 0, err
	}
	return config.Width, config.Height, nil
}

func (p WEBPInfo) ExtractInfo(data []byte) (width, height int, err error) {
	config, err := webp.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return 0, 0, err
	}
	return config.Width, config.Height, nil
}
