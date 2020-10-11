package model

type HTMLStructure struct {
	HTMLVersion           string
	PageTitle             string
	HeadingCount          HeadingStructure
	InternalLinkCount     int64
	ExternalLinkCount     int64
	InaccessibleLinkCount int64
	LoginFormExist        bool
}
