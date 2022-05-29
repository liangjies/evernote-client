package service

import (
	"evernote-client/global"
	"time"
)

//@function: GetRedis
//@description: 从Redis取数据
//@param: prefix + key string
//@return: err error, redisJWT string
func GetRedis(key string) (err error, value string) {
	prefix := global.SYS_CONFIG.Redis.Prefix
	value, err = global.SYS_REDIS.Get(prefix + key).Result()
	return err, value
}

//@function: SetRedis
//@description: 存入Redis并设置过期时间
//@param: key string, value string, expTime uint
//@return: err error
func SetRedis(key string, value string, expTime uint) (err error) {
	prefix := global.SYS_CONFIG.Redis.Prefix
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(expTime) * time.Second
	err = global.SYS_REDIS.Set(prefix+key, value, timer).Err()
	return err
}

//@function: DelRedis
//@description: 删除Redis数据
//@param: prefix + key string
//@return: err error
func DelRedis(key string) (err error) {
	prefix := global.SYS_CONFIG.Redis.Prefix
	err = global.SYS_REDIS.Del(prefix + key).Err()
	return err
}
