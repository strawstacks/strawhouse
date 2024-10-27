package strawhouse

type SignatureMode uint8

const (
	SignatureModeFile SignatureMode = iota
	SignatureModeDirectory
)

type SignatureAction uint8

const (
	SignatureActionGet SignatureAction = iota
	SignatureActionUpload
)

type FeedType uint8

const (
	FeedTypeGet FeedType = iota
	FeedTypeUpload
	FeedTypeDelete
)
