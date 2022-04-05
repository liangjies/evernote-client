package service

import (
	"evernote-client/global"
	"evernote-client/model"
	"evernote-client/utils/upload"
	"mime/multipart"
	"strings"
)

//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func Upload(file model.FileUpload) error {
	return global.SYS_DB.Create(&file).Error
}

//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func UploadFile(header *multipart.FileHeader, noSave string) (err error, file model.FileUpload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.FileUpload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return Upload(f), f
	}
	return
}
