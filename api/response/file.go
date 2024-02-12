package response

type (
	UploadResult struct {
		FileID string `json:"file_id"`
		Link   string `json:"link"`
	}
)
