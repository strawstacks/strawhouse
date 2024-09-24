package payload

type FileInfo struct {
	Filename  *string `json:"filename"`
	Directory *string `json:"directory"`
}
