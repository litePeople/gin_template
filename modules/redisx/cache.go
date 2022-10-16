package redisx

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Cache interface {
	Set(string, string) error
	SetExp(string, string, time.Duration) error
	SetNx(string, string, time.Duration) error
	DefaultExp() time.Duration
	GetSet(string, string) (string, error)
	Get(string) (string, error)
	Del(string) error
	ExistKey(string) (bool, error)
	GetListLast(string) (string, error)
	ListRPush(string interface{}) (int64, error)
	LenList(string) (int64, error)
	HashExistKey(string) (bool error)
	HashSet(string) error
	HashGet(string) (string error)
	HashDelField(string) error
}

type client struct {
	cli *redis.Client `ini:"-"`
	// 地址
	Host string
	// 端口
	Port uint
	// 数据库
	Db uint
	// 密码
	Password string
}

// Expire 设置 key的过期时间
func (c *client) Expire(key string, value time.Duration) error {
	statusCmd := c.cli.Expire(key, value)
	return statusCmd.Err()
}

func (c *client) Ttl(key string) (float64, error) {
	intCmd, err := c.cli.TTL(key).Result()
	return intCmd.Seconds(), err
}
func (c *client) Set(key string, value string) (err error) {
	statusCmd := c.cli.Set(key, value, c.DefaultExp())
	return statusCmd.Err()

}

func (c *client) ExistKey(key string) (bool, error) {
	boolCmd := c.cli.Exists(key)
	return boolCmd.Val() > 0, boolCmd.Err()
}

func (c *client) DefaultExp() time.Duration {
	return time.Second * 10
}

func (c *client) SetExp(key string, value string, exp time.Duration) (err error) {
	statusCmd := c.cli.Set(key, value, exp)
	return statusCmd.Err()
}

func (c *client) GetSet(key string, value string) (oldValue string, err error) {
	stringCmd := c.cli.GetSet(key, value)
	return stringCmd.Val(), stringCmd.Err()
}

func (c *client) Get(key string) (value string, err error) {
	stringCmd := c.cli.Get(key)
	return stringCmd.Val(), stringCmd.Err()
}

func (c *client) Del(key string) error {
	intCmd := c.cli.Del(key)
	return intCmd.Err()
}

// IncrBy 增量
func (c *client) IncrBy(key string, value int64) error {
	statusCmd := c.cli.IncrBy(key, value)
	return statusCmd.Err()
}

// DecrBy 减量
func (c *client) DecrBy(key string, value int64) error {
	statusCmd := c.cli.DecrBy(key, value)
	return statusCmd.Err()
}

// GetListLast 从队列弹出一个
func (c *client) GetListLast(key string) (value string, err error) {
	stringCmd := c.cli.LPop(key)
	return stringCmd.Val(), stringCmd.Err()
}

// LenList 获取队列得长度
func (c *client) LenList(key string) (value int64, err error) {
	intCmd := c.cli.LLen(key)
	return intCmd.Val(), intCmd.Err()
}

// ListRPush push到队列
func (c *client) ListRPush(key string, value interface{}) (v int64, err error) {
	intCmd := c.cli.RPush(key, value)
	return intCmd.Val(), intCmd.Err()
}

func (c *client) HashExistKey(key, name string) (value bool, err error) {
	boolCmd := c.cli.HExists(key, name)
	return boolCmd.Val(), boolCmd.Err()
}
func (c *client) HashSet(key string, field map[string]interface{}) error {
	statusCmd := c.cli.HMSet(key, field)
	return statusCmd.Err()
}
func (c *client) HashGet(key string, field string) (value string, err error) {
	strCmd := c.cli.HGet(key, field)
	return strCmd.Val(), strCmd.Err()
}
func (c *client) HashDelField(key string, field string) (err error) {
	statusCmd := c.cli.HDel(key, field)
	return statusCmd.Err()
}

func (c *client) HashGetAll(key string) (value map[string]string, err error) {
	strCmd := c.cli.HGetAll(key)
	return strCmd.Val(), strCmd.Err()
}

// Hincrby 往指定的key增加指定的int的值
func (c *client) Hincrby(key, field string, num int64) error {
	structCmd := c.cli.HIncrBy(key, field, num)
	return structCmd.Err()
}

// HincrbyFloat 往指定的key值增加指定float64的值
func (c *client) HincrbyFloat(key, field string, num float64) error {
	structCmd := c.cli.HIncrByFloat(key, field, num)
	return structCmd.Err()
}

// ZrevRange 有序集合
func (c *client) ZrevRange(key string, start, stop int64) ([]string, error) {
	valueCmd := c.cli.ZRevRange(key, start, stop)
	return valueCmd.Val(), valueCmd.Err()
}
func (c *client) ZRevRangeWithScores(key string, start, stop int64) ([]map[string]interface{}, error) {
	valueCmd := c.cli.ZRevRangeWithScores(key, start, stop)
	if valueCmd.Err() != nil {
		return nil, valueCmd.Err()
	}
	resultes := make([]map[string]interface{}, 0, len(valueCmd.Val()))
	for _, redis := range valueCmd.Val() {
		result := make(map[string]interface{})
		result["member"] = redis.Member
		result["score"] = redis.Score
		resultes = append(resultes, result)
	}
	return resultes, nil
}

func (c *client) ZRevRangeByScores(key string, start, stop int64) ([]map[string]interface{}, error) {
	zrange := redis.ZRangeBy{
		Min: fmt.Sprintf("%d", start),
		Max: fmt.Sprintf("%d", stop),
	}
	valueCmd := c.cli.ZRevRangeByScoreWithScores(key, zrange)
	if valueCmd.Err() != nil {
		return nil, valueCmd.Err()
	}
	resultes := make([]map[string]interface{}, 0, len(valueCmd.Val()))
	for _, redis := range valueCmd.Val() {
		result := make(map[string]interface{})
		result["member"] = redis.Member
		result["score"] = redis.Score
		resultes = append(resultes, result)
	}
	return resultes, nil
}

// Zincrby 增加一个成员的分数
func (c *client) Zincrby(key string, score float64, member string) error {
	statusCmd := c.cli.ZIncrBy(key, score, member)
	return statusCmd.Err()
}

// ZAdd 往集合里面设置一个成员
func (c *client) ZAdd(key string, score float64, member string) error {
	z := redis.Z{Member: member, Score: score}
	statusCmd := c.cli.ZAdd(key, z)
	return statusCmd.Err()
}

// ZRem 移除集合里面一个成员或多个成员
func (c *client) ZRem(keys string, members []string) error {
	interfaces := make([]interface{}, 0, len(members))
	for _, member := range members {
		interfaces = append(interfaces, member)
	}
	intCmd := c.cli.ZRem(keys, interfaces...)
	return intCmd.Err()
}

// SAdd 往无需集合设置一个成员
func (c *client) SAdd(key string, member string) error {
	return c.cli.SAdd(key, member).Err()
}

// SCard 获取无序集合的成员总数
func (c *client) SCard(key string) (int64, error) {
	resCmd := c.cli.SCard(key)
	return resCmd.Val(), resCmd.Err()
}

// SMembers 获取无序集合的所有成员
func (c *client) SMembers(key string) ([]string, error) {
	resCmd := c.cli.SMembers(key)
	return resCmd.Val(), resCmd.Err()
}

// SRem 删除无语集合的成员
func (c *client) SRem(key string, member ...string) error {
	return c.cli.SRem(key, member).Err()
}

// SetNx 设置一个redis锁
func (c *client) SetNx(key, val string, exp time.Duration) error {
	return c.cli.SetNX(key, val, exp).Err()
}

// Lock 锁函数
// lockKey redis key
// lockListTime 锁的存活时间
// checkTime 尝试获取锁间隔时间
// callback 获取锁后的操作
func (c *client) Lock(lockKey string, lockLifeTime time.Duration, checkTime time.Duration, callback func()) error {
	// 上锁
	if err := c.SetNx(lockKey, "", lockLifeTime); err != nil { // 上锁失败，代表已经有程序在执行中
		time.Sleep(checkTime)
		return c.Lock(lockKey, lockLifeTime, checkTime, callback)
	}

	// 解锁
	defer c.Del(lockKey)

	callback()
	return nil
}
