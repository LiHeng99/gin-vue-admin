package model

import (
	"gin-vue-admin/global"
	"github.com/jinzhu/gorm"
)

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text"`
}

// @title    JsonInBlacklist
// @description   create jwt blacklist
// @auth                     （2020/04/05  20:22 ）
// @return    err              error
func (j *JwtBlacklist) JsonInBlacklist() (err error) {
	err = global.GVA_DB.Create(j).Error
	return
}

// @title    IsBlacklist
// @description   check if the Jwt is in the blacklist or not, 判断JWT是否在黑名单内部
// @auth                     （2020/04/05  20:22 ）
// @param     newPassword     string
// @return    err             error
func (j *JwtBlacklist) IsBlacklist(Jwt string) bool {
	isNotFound := global.GVA_DB.Where("jwt = ?", Jwt).First(j).RecordNotFound()
	return !isNotFound
}

// @title    GetRedisJWT
// @description   Get user info in redis
// @auth                     （2020/04/05  20:22 ）
// @param     newPassword     string
// @return    err             error
func (j *JwtBlacklist) GetRedisJWT(userName string) (err error, RedisJWT string) {
	RedisJWT, err = global.GVA_REDIS.Get(userName).Result()
	return err, RedisJWT
}

// @title    SetRedisJWT
// @description   set jwt into the Redis
// @auth                     （2020/04/05  20:22 ）
// @param     userName         string
// @return    err              error
func (j *JwtBlacklist) SetRedisJWT(userName string) (err error) {
	err = global.GVA_REDIS.Set(userName, j.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
