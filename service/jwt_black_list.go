package service

import (
	"errors"
	"evernote-client/global"
	"evernote-client/model"
	"time"

	"gorm.io/gorm"
)

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.SYS_DB.Create(&jwtList).Error
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	err := global.SYS_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: prefix + userName string
//@return: err error, redisJWT string

func GetRedisJWT(userName string) (err error, redisJWT string) {
	prefix := global.SYS_CONFIG.Redis.Prefix
	redisJWT, err = global.SYS_REDIS.Get(prefix + userName).Result()
	return err, redisJWT
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func SetRedisJWT(jwt string, userName string) (err error) {
	prefix := global.SYS_CONFIG.Redis.Prefix
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.SYS_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.SYS_REDIS.Set(prefix+userName, jwt, timer).Err()
	return err
}
