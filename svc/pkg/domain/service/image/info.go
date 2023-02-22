package image

type InfoExtractor interface {
	ExtractInfo(data []byte) (width, height int, err error)
}
