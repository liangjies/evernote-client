package service

import (
	"context"
	"evernote-client/global"
	"time"
)

// @function: GetRedis
// @description: 从Redis取数据
// @param: prefix + key string
// @return: err error, redisJWT string
func GetRedis(key string) (err error, value string) {
	prefix := global.CONFIG.Redis.Prefix
	ctx := context.Background()
	value, err = global.REDIS.Get(ctx, prefix+key).Result()
	return err, value
}

// @function: SetRedis
// @description: 存入Redis并设置过期时间
// @param: key string, value string, expTime uint
// @return: err error
func SetRedis(key string, value string, expTime uint) (err error) {
	prefix := global.CONFIG.Redis.Prefix
	ctx := context.Background()
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(expTime) * time.Second
	err = global.REDIS.Set(ctx, prefix+key, value, timer).Err()
	return err
}

// @function: DelRedis
// @description: 删除Redis数据
// @param: prefix + key string
// @return: err error
func DelRedis(key string) (err error) {
	prefix := global.CONFIG.Redis.Prefix
	ctx := context.Background()
	err = global.REDIS.Del(ctx, prefix+key).Err()
	return err
}
