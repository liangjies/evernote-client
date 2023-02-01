package response

import "evernote-client/model"

type FileUploadResponse struct {
	File model.EvnUpload `json:"file"`
}
