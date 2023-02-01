package upload

import (
	"evernote-client/global"
	"mime/multipart"
)

//@interface_name: OSS
//@description: OSS接口

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// @function: NewOss
// @description: OSS接口
// @description: OSS的实例化方法
// @return: OSS
func NewOss() OSS {
	switch global.CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "qiniu":
		return &Qiniu{}
	case "tencent-cos":
		return &TencentCOS{}
	default:
		return &Local{}
	}
}
