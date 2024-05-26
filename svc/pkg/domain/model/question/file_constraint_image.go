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
		Ratio         RatioSpec
		MinNumber     *int
		MaxNumber     *int
		Width, Height DimensionSpec
		PNGInfo       image.InfoExtractor
		JPGInfo       image.InfoExtractor
		WEBPInfo      image.InfoExtractor
	}
	ImageType int
)

const (
	PNG  ImageType = 1
	JPG  ImageType = 2
	WEBP ImageType = 3
	TIFF ImageType = 4
	HEIC ImageType = 5
	SVG  ImageType = 6
)

func NewImageFileConstraint(
	minNumber, maxNumber *int,
	width, height DimensionSpec,
	ratio RatioSpec,
) (*ImageFileConstraint, error) {
	if err := width.Validate(); err != nil {
		return nil, fmt.Errorf("invalid width: %w", err)
	}
	if err := height.Validate(); err != nil {
		return nil, fmt.Errorf("invalid height: %w", err)
	}
	if err := ratio.Validate(); err != nil {
		return nil, fmt.Errorf("invalid ratio: %w", err)
	}
	return &ImageFileConstraint{
		Ratio:     ratio,
		MinNumber: minNumber,
		MaxNumber: maxNumber,
		Width:     width,
		Height:    height,
		PNGInfo:   imagePkg.NewPNGInfo(),
		JPGInfo:   imagePkg.NewJPEGInfo(),
		WEBPInfo:  imagePkg.NewWEBPInfo(),
	}, nil
}

func (c ImageFileConstraint) GetFileType() FileType {
	return Image
}

func (c ImageFileConstraint) GetExtensions() []Extension {
	return []Extension{".jpg", ".jpeg", ".png", ".webp"}
}

func (c ImageFileConstraint) ValidateFiles(files []File) error {
	if len(files) == 0 {
		return errors.New("no file found")
	}
	if c.MinNumber != nil && *c.MinNumber <= len(files) {
		return errors.New(fmt.Sprintf(
			"number of files not satisfied. min number: %d, actual number: %d", *c.MinNumber, len(files)))
	}
	if c.MaxNumber != nil && *c.MaxNumber >= len(files) {
		return errors.New(fmt.Sprintf(
			"number of files not satisfied. max number: %d, actual number: %d", *c.MaxNumber, len(files)))
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
	if width <= 0 || height <= 0 {
		return errors.New("invalid image size")
	}
	if err := validateDimension(c.Width, width, "width"); err != nil {
		return err
	}
	if err := validateDimension(c.Height, height, "height"); err != nil {
		return err
	}
	if err := validateRatio(c.Ratio, float32(width)/float32(height), "ratio"); err != nil {
		return err
	}
	return nil
}

func validateDimension(d DimensionSpec, value int, name string) error {
	if d.Min != nil && value < *d.Min {
		return errors.New(
			fmt.Sprintf("%s not satisfied. min %s: %d, actual %s: %d", name, name, *d.Min, name, value))
	}
	if d.Max != nil && value > *d.Max {
		return errors.New(
			fmt.Sprintf("%s not satisfied. max %s: %d, actual %s: %d", name, name, *d.Max, name, value))
	}
	if d.Eq != nil && value != *d.Eq {
		return errors.New(
			fmt.Sprintf("%s not satisfied. expected %s: %d, actual %s: %d", name, name, *d.Eq, name, value))
	}
	return nil
}

func validateRatio(r RatioSpec, value float32, name string) error {
	if r.Min != nil && value < *r.Min {
		return errors.New(
			fmt.Sprintf("%s not satisfied. min %s: %f, actual %s: %f", name, name, *r.Min, name, value))
	}
	if r.Max != nil && value > *r.Max {
		return errors.New(
			fmt.Sprintf("%s not satisfied. max %s: %f, actual %s: %f", name, name, *r.Max, name, value))
	}
	if r.Eq != nil && value != *r.Eq {
		return errors.New(
			fmt.Sprintf("%s not satisfied. expected %s: %f, actual %s: %f", name, name, *r.Eq, name, value))
	}
	return nil
}

func (c ImageFileConstraint) checkExtension(ext string) (ImageType, error) {
	for _, e := range c.GetExtensions() {
		if string(e) == ext {
			return convertToImageType(ext)
		}
	}
	return 0, errors.New(
		fmt.Sprintf("invalid file type. specified extensions: %v", c.GetExtensions()))
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

func (s RatioSpec) Validate() error {
	if s.Eq != nil {
		if *s.Eq <= 0 {
			return errors.New("eq ratio must be positive")
		}
		if s.Min != nil || s.Max != nil {
			return errors.New("eq ratio is specified with min or max ratio")
		}
		return nil
	}

	if s.Min != nil && *s.Min <= 0 {
		return errors.New("min ratio must be positive")
	}
	if s.Max != nil && *s.Max <= 0 {
		return errors.New("max ratio must be positive")
	}
	if s.Min != nil && s.Max != nil && *s.Min > *s.Max {
		return errors.New("min ratio is greater than max ratio")
	}
	return nil
}

func (s DimensionSpec) Validate() error {
	if s.Eq != nil {
		if *s.Eq <= 0 {
			return errors.New("eq dimension must be positive")
		}
		if s.Min != nil || s.Max != nil {
			return errors.New("eq dimension is specified with min or max dimension")
		}
		return nil
	}
	if s.Min != nil && *s.Min <= 0 {
		return errors.New("min dimension must be positive")
	}
	if s.Max != nil && *s.Max <= 0 {
		return errors.New("max dimension must be positive")
	}
	return nil
}

const (
	FileImageConstraintWidth        = "w"
	FileImageConstraintHeight       = "h"
	FileImageConstraintRatio        = "r"
	FileImageConstraintMinNumber    = "min"
	FileImageConstraintMaxNumber    = "max"
	FileImageConstraintDimensionEq  = "eq"
	FileImageConstraintDimensionMin = "min"
	FileImageConstraintDimensionMax = "max"
	FileImageConstraintRatioEq      = "eq"
	FileImageConstraintRatioMin     = "min"
	FileImageConstraintRatioMax     = "max"
)

func ImportImageFileConstraint(c map[string]interface{}) (*ImageFileConstraint, error) {
	width, err := loadDimensionSpec(c, FileImageConstraintWidth)
	if err != nil {
		return nil, fmt.Errorf("failed to load width: %w", err)
	}
	height, err := loadDimensionSpec(c, FileImageConstraintHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to load height: %w", err)
	}
	ratio, err := loadRatioSpec(c, FileImageConstraintRatio)
	if err != nil {
		return nil, fmt.Errorf("failed to load ratio: %w", err)
	}
	minNumber, err := loadInt(c, FileImageConstraintMinNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to load min number: %w", err)
	}
	maxNumber, err := loadInt(c, FileImageConstraintMaxNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to load max number: %w", err)
	}
	return NewImageFileConstraint(minNumber, maxNumber, width, height, ratio)
}

func loadInt(t map[string]interface{}, key string) (*int, error) {
	v, has := t[key]
	if !has {
		return nil, nil
	}
	i, ok := v.(int)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid %s", key))
	}
	return &i, nil
}

func (c ImageFileConstraint) Export() map[string]interface{} {
	result := map[string]interface{}{}
	if c.MinNumber != nil {
		result[FileImageConstraintMinNumber] = *c.MinNumber
	}
	if c.MaxNumber != nil {
		result[FileImageConstraintMaxNumber] = *c.MaxNumber
	}
	widthC := c.Width.Export()
	result[FileImageConstraintWidth] = widthC
	heightC := c.Height.Export()
	result[FileImageConstraintHeight] = heightC
	return result
}
