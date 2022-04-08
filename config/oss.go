package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`            // 存储区域
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`      // 空间名称
	ImgPath       string `mapstructure:"img-path" json:"imgPath" yaml:"img-path"` // CDN加速域名
	PathPrefix    string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`                  // 是否使用https
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`               // 秘钥AK
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`               // 秘钥SK
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"` // 上传是否使用CDN上传加速
}
type TencentCOS struct {
	Bucket     string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Region     string `mapstructure:"region" json:"region" yaml:"region"`
	SecretID   string `mapstructure:"secret-id" json:"secretID" yaml:"secret-id"`
	SecretKey  string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	BaseURL    string `mapstructure:"base-url" json:"baseURL" yaml:"base-url"`
	PathPrefix string `mapstructure:"path-prefix" json:"pathPrefix" yaml:"path-prefix"`
}
