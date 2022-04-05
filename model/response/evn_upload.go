package response

import "evernote-client/model"

type FileUploadResponse struct {
	File model.FileUpload `json:"file"`
}
