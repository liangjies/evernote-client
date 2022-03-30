package config

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// gorm
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
}
