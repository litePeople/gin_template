package redisx

import (
	"github.com/go-redis/redis"
	"time"
)

func GetClient() *redis.Client {
	return cache.cli
}

// Expire 设置 key的过期时间
func Expire(key string, value time.Duration) error {
	return cache.Expire(key, value)
}

func Ttl(key string) (float64, error) {
	return cache.Ttl(key)
}
func Set(key string, value string) (err error) {
	return cache.Set(key, value)

}

func ExistKey(key string) (bool, error) {
	return cache.ExistKey(key)
}

func DefaultExp() time.Duration {
	return cache.DefaultExp()
}

func SetExp(key string, value string, exp time.Duration) (err error) {
	return cache.SetExp(key, value, exp)
}

func GetSet(key string, value string) (oldValue string, err error) {
	return cache.GetSet(key, value)
}

func Get(key string) (value string, err error) {
	return cache.Get(key)
}

func Del(key string) error {
	return cache.Del(key)
}

// IncrBy 增量
func IncrBy(key string, value int64) error {
	return cache.IncrBy(key, value)
}

// DecrBy 减量
func DecrBy(key string, value int64) error {
	return cache.DecrBy(key, value)
}

// GetListLast 从队列弹出一个
func GetListLast(key string) (value string, err error) {
	return cache.GetListLast(key)
}

// LenList 获取队列得长度
func LenList(key string) (value int64, err error) {
	return cache.LenList(key)
}

// ListRPush push到队列
func ListRPush(key string, value interface{}) (v int64, err error) {
	return cache.ListRPush(key, value)
}

func HashExistKey(key, name string) (value bool, err error) {
	return cache.HashExistKey(key, name)
}
func HashSet(key string, field map[string]interface{}) error {
	return cache.HashSet(key, field)
}
func HashGet(key string, field string) (value string, err error) {
	return cache.HashGet(key, field)
}
func HashDelField(key string, field string) (err error) {
	return cache.HashDelField(key, field)
}

func HashGetAll(key string) (value map[string]string, err error) {
	return cache.HashGetAll(key)
}

// Hincrby 往指定的key增加指定的int的值
func Hincrby(key, field string, num int64) error {
	return cache.Hincrby(key, field, num)
}

// HincrbyFloat 往指定的key值增加指定float64的值
func HincrbyFloat(key, field string, num float64) error {
	return cache.HincrbyFloat(key, field, num)
}

// ZrevRange 有序集合
func ZrevRange(key string, start, stop int64) ([]string, error) {
	return cache.ZrevRange(key, start, stop)
}
func ZRevRangeWithScores(key string, start, stop int64) ([]map[string]interface{}, error) {
	return cache.ZRevRangeWithScores(key, start, stop)
}

func ZRevRangeByScores(key string, start, stop int64) ([]map[string]interface{}, error) {
	return cache.ZRevRangeByScores(key, stop, stop)
}

// Zincrby 增加一个成员的分数
func Zincrby(key string, score float64, member string) error {
	return cache.Zincrby(key, score, member)
}

// ZAdd 往集合里面设置一个成员
func ZAdd(key string, score float64, member string) error {
	return cache.ZAdd(key, score, member)
}

// ZRem移除集合里面一个成员或多个成员
func ZRem(keys string, members []string) error {
	return cache.ZRem(keys, members)
}

// SAdd 往无需集合设置一个成员
func SAdd(key string, member string) error {
	return cache.SAdd(key, member)
}

// SCard 获取无序集合的成员总数
func SCard(key string) (int64, error) {
	return cache.SCard(key)
}

// SMembers 获取无序集合的所有成员
func SMembers(key string) ([]string, error) {
	return cache.SMembers(key)
}

// SRem 删除无语集合的成员
func SRem(key string, member ...string) error {
	return cache.SRem(key, member...)
}

// SetNx 设置一个redis锁
func SetNx(key, val string, exp time.Duration) error {
	return cache.SetNx(key, val, exp)
}

// Lock 锁函数
// lockKey redis key
// lockListTime 锁的存活时间
// checkTime 尝试获取锁间隔时间
// callback 获取锁后的操作
func Lock(lockKey string, lockLifeTime time.Duration, checkTime time.Duration, callback func()) error {
	return cache.Lock(lockKey, lockLifeTime, checkTime, callback)
}
