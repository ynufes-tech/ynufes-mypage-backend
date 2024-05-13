package question

import (
	"errors"
	"fmt"
	"path/filepath"
	imagePkg "ynufes-mypage-backend/pkg/image"
	"ynufes-mypage-backend/svc/pkg/domain/service/image"
)

type (
	ImageFileConstraint struct {
		Ratio               float64
		MinNumber           int
		MaxNumber           int
		MinResolutionWidth  int
		MaxResolutionWidth  int
		MinResolutionHeight int
		MaxResolutionHeight int
		Extensions          []Extension
		PNGInfo             image.InfoExtractor
		JPGInfo             image.InfoExtractor
		WEBPInfo            image.InfoExtractor
	}
	ImageType int
)

const (
	PNG  ImageType = 1
	JPG  ImageType = 2
	WEBP ImageType = 3
)

func NewImageFileConstraint(
	ratio float64, minNumber, maxNumber, minWidth, maxWidth, minHeight, maxHeight int, extensions []Extension,
) ImageFileConstraint {
	return ImageFileConstraint{
		Ratio:               ratio,
		MinNumber:           minNumber,
		MaxNumber:           maxNumber,
		MinResolutionWidth:  minWidth,
		MaxResolutionWidth:  maxWidth,
		MinResolutionHeight: minHeight,
		MaxResolutionHeight: maxHeight,
		Extensions:          extensions,
		PNGInfo:             imagePkg.NewPNGInfo(),
		JPGInfo:             imagePkg.NewJPEGInfo(),
		WEBPInfo:            imagePkg.NewWEBPInfo(),
	}
}

func ImportImageFileConstraint(standard StandardFileConstraint) ImageFileConstraint {
	ratio, _ := standard.Customs["ratio"].(float64)
	minNumber, _ := standard.Customs["minNumber"].(int64)
	maxNumber, _ := standard.Customs["maxNumber"].(int64)
	minWidth, _ := standard.Customs["minWidth"].(int64)
	maxWidth, _ := standard.Customs["maxWidth"].(int64)
	minHeight, _ := standard.Customs["minHeight"].(int64)
	maxHeight, _ := standard.Customs["maxHeight"].(int64)
	extsI, _ := standard.Customs["extensions"].([]interface{})
	exts := make([]Extension, len(extsI))
	for i, extI := range extsI {
		exts[i] = extI.(Extension)
	}

	return NewImageFileConstraint(
		ratio, int(minNumber), int(maxNumber), int(minWidth), int(maxWidth), int(minHeight), int(maxHeight), exts)
}

func (c ImageFileConstraint) Export() (*StandardFileConstraint, error) {
	return NewStandardFileConstraint(Image,
		map[string]interface{}{
			"ratio":      c.Ratio,
			"minNumber":  c.MinNumber,
			"maxNumber":  c.MaxNumber,
			"minWidth":   c.MinResolutionWidth,
			"maxWidth":   c.MaxResolutionWidth,
			"minHeight":  c.MinResolutionHeight,
			"maxHeight":  c.MaxResolutionHeight,
			"extensions": c.Extensions,
		})
}

func (c ImageFileConstraint) GetFileType() FileType {
	return Image
}

func (c ImageFileConstraint) GetExtensions() []Extension {
	return c.Extensions
}

func (c ImageFileConstraint) ValidateFiles(files []File) error {
	if c.MinNumber > 0 && len(files) < c.MinNumber {
		return errors.New(fmt.Sprintf(
			"number of files not satisfied. min number: %d, actual number: %d", c.MinNumber, len(files)))
	}
	if c.MaxNumber > 0 && len(files) > c.MaxNumber {
		return errors.New(fmt.Sprintf(
			"number of files not satisfied. max number: %d, actual number: %d", c.MaxNumber, len(files)))
	}

	for _, file := range files {
		// get file extension
		ext := filepath.Ext(file.FileName)
		imgType, err := c.checkExtension(ext)
		if err != nil {
			return err
		}
		err = c.validateProperties(imgType, file.Data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c ImageFileConstraint) validateProperties(imgType ImageType, file []byte) error {
	var width, height int
	var err error
	switch imgType {
	case PNG:
		width, height, err = c.PNGInfo.ExtractInfo(file)
	case JPG:
		width, height, err = c.JPGInfo.ExtractInfo(file)
	case WEBP:
		width, height, err = c.WEBPInfo.ExtractInfo(file)
	default:
		// skip check for unimplemented image type
		return nil
	}
	if err != nil {
		return err
	}
	if c.MinResolutionWidth > 0 && width < c.MinResolutionWidth {
		return errors.New(
			fmt.Sprintf("width not satisfied. min width: %d, actual width: %d", c.MinResolutionWidth, width))
	}
	if c.MaxResolutionWidth > 0 && width > c.MaxResolutionWidth {
		return errors.New(
			fmt.Sprintf("width not satisfied. max width: %d, actual width: %d", c.MaxResolutionWidth, width))
	}
	if c.MinResolutionHeight > 0 && height < c.MinResolutionHeight {
		return errors.New(
			fmt.Sprintf("height not satisfied. min height: %d, actual height: %d", c.MinResolutionHeight, height))
	}
	if c.MaxResolutionHeight > 0 && height > c.MaxResolutionHeight {
		return errors.New(
			fmt.Sprintf("height not satisfied. max height: %d, actual height: %d", c.MaxResolutionHeight, height))
	}
	if c.Ratio > 0 {
		if ratio := float64(width) / float64(height); ratio != c.Ratio {
			return errors.New(
				fmt.Sprintf("ratio not satisfied. expected ratio: %f, actual ratio: %f", c.Ratio, ratio))
		}
	}
	return nil
}

func (c ImageFileConstraint) checkExtension(ext string) (ImageType, error) {
	if len(c.Extensions) == 0 {
		// if extension is not specified, check with default extensions
		return convertToImageType(ext)
	}
	for _, e := range c.Extensions {
		if string(e) == ext {
			return convertToImageType(ext)
		}
	}
	return 0, errors.New(
		fmt.Sprintf("invalid file type. specified extensions: %v", c.Extensions))
}

func convertToImageType(ext string) (ImageType, error) {
	switch ext {
	case ".jpg", ".jpeg":
		return JPG, nil
	case ".png":
		return PNG, nil
	case ".webp":
		return WEBP, nil
	default:
		return 0, errors.New(
			fmt.Sprintf("invalid file type. available extensions: %v",
				[]string{".jpg", ".jpeg", ".png", ".webp"}))
	}
}
