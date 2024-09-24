package payload

type UploadResponse struct {
	Path *string `json:"path"`
	Hash *string `json:"hash"`
	Size *int64  `json:"size"`
}
