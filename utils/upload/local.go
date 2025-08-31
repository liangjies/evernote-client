package upload

import (
	"errors"
	"evernote-client/global"
	"evernote-client/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 简单大小限制：默认 10MB
	const maxSize = 10 * 1024 * 1024
	if file.Size > maxSize {
		return "", "", errors.New("file too large: max 10MB")
	}
	// 读取并校验后缀与 MIME（弱校验，强校验建议读取魔数或专门库）
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	lower := strings.ToLower(ext)
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".bmp": true, ".webp": true, ".svg": false, ".pdf": true, ".txt": true, ".md": true}
	if ok, exist := allowedExt[lower]; !exist || !ok {
		return "", "", errors.New("unsupported file type")
	}
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.CONFIG.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	// 通过 content-type 做一层额外校验（非强保证）
	// 注意：multipart 提供的 Header 可能被伪造，必要时读取前 512 字节 sniff
	if ct := file.Header.Get("Content-Type"); ct != "" {
		// 允许的通用类型
		if !strings.HasPrefix(ct, "image/") && ct != "application/pdf" && ct != "text/plain" && ct != "text/markdown" {
			// 如果 Header 没给或给错，尝试 sniff 一下
			buf := make([]byte, 512)
			n, _ := f.Read(buf)
			_, _ = f.Seek(0, 0)
			sniff := http.DetectContentType(buf[:n])
			if !strings.HasPrefix(sniff, "image/") && sniff != "application/pdf" && sniff != "text/plain" {
				return "", "", errors.New("invalid content type")
			}
		}
	}

	out, createErr := os.Create(p)
	if createErr != nil {
		global.LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	p := global.CONFIG.Local.Path + "/" + key
	if strings.Contains(p, global.CONFIG.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
