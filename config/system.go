package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 端口值
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis" json:"useRedis" yaml:"use-redis"`                // 多点登录拦截
	VersionMax    uint   `mapstructure:"version-max" json:"versionMax" yaml:"version-max"`          // 笔记历史记录保存最多数量
}
