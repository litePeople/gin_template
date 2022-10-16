package mongox

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	cli *mongo.Client `ini:"-"`
	// 地址
	Host string
	// 端口
	Port uint
	// 名称
	Name string
	// 用户名
	User string
	// 密码
	Password string
}

func GetClient() *client {
	return &mgo
}

// DB 切换数据库
// cli 客户端
// *mongo.Database 返回mongo的数据库
func (mc *client) DB(client *mongo.Client) *mongo.Database {
	return client.Database(mc.Name)
}

func (mc *client) WithC(collection string, job func(*mongo.Collection) error) error {
	return job(mc.DB(mc.cli).Collection(collection))
}

// Upsert 存在则更新，不存在则创建
// collection 集合名称
// selector 筛选条件
// change 文档的数据
// error 错误信息
func (mc *client) Upsert(collection string, selector interface{}, change interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		var upsert = true
		_, err := c.UpdateOne(context.Background(), selector, change, &options.UpdateOptions{Upsert: &upsert})
		return err
	})
}

// UpsertById 根据id插入或更新
// collection 集合名称
// id id
// update 文档的数据
// error 错误信息
func (mc *client) UpsertById(collection string, id interface{}, update interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		var upsert = true
		_, err := c.UpdateByID(context.Background(), id, update, &options.UpdateOptions{Upsert: &upsert})
		return err
	})
}

// UpdateById 根据id更新，当id不存在的时候不会报错
// collection 集合名称
// id id
// change 文档的数据
// error 错误信息
func (mc *client) UpdateById(collection string, id interface{}, change interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		_, err := c.UpdateByID(context.Background(), id, change)
		return err
	})
}

// UpdateOne 根据查询条件更新一条记录
// collection 集合名称
// selector 查询条件
// change 文档的数据
// error 错误信息
func (mc *client) UpdateOne(collection string, selector, change interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		_, err := c.UpdateOne(context.Background(), selector, change)
		return err
	})
}

// UpdateAll 根据查询条件更新所有符合条件的记录
// collection 集合名称
// selector 查询条件
// change 文档数据
// error 错误信息
func (mc *client) UpdateAll(collection string, selector, change interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		_, err := c.UpdateMany(context.Background(), selector, change)
		return err
	})
}

// Insert 批量插入文档
// collection 集合名称
// data 数据
// error 错误信息
func (mc *client) Insert(collection string, data ...interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		_, err := c.InsertMany(context.Background(), data)
		return err
	})
}

// All 根据查询条件获取所有符合条件的记录
// collection 集合名称
// query 查询条件
// result 查询结果
// error 错误信息
func (mc *client) All(collection string, query interface{}, result interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		cursor, err := c.Find(context.Background(), query)
		if nil != err {
			return err
		}

		return cursor.All(context.Background(), result)
	})
}

// AllSelect 根据查询条件查询所有符合条件的记录，并且限制字段
// collection 集合名称
// query 查询条件
// projection 限制字段
// result 查询结果
// error 错误信息
func (mc *client) AllSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		cursor, err := c.Find(context.Background(), query, options.Find().SetProjection(projection))
		if nil != err {
			return err
		}
		return cursor.All(context.Background(), result)
	})
}

// One 根据查询条件查找一条记录
// collection 集合名称
// query 查询条件
// result 查询结果
// error 错误信息
func (mc *client) One(collection string, query interface{}, result interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		return c.FindOne(context.Background(), query).Decode(result)
	})
}

// OneSelect 查找一条记录，并且筛选指定的字段
// collection 集合名
// query 查询条件
// projection 限制字段
// result 查询结果
// error 错误信息
func (mc *client) OneSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		return c.FindOne(context.Background(), query, options.FindOne().SetProjection(projection)).Decode(result)
	})
}

// FindById 根据id获取文档记录
// collection 集合名称
// id id
// result 查询结果
// error 错误信息
func (mc *client) FindById(collection string, id interface{}, result interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		return c.FindOne(context.Background(), bson.M{"_id": id}).Decode(result)
	})
}

// RemoveById 根据id删除文档
// collection 集合名称
// id id
// error 错误信息
func (mc *client) RemoveById(collection string, id interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		_, err := c.DeleteOne(context.Background(), bson.M{"_id": id})
		return err
	})
}

// RemoveOne 根据条件删除一条记录
// collection 集合名称
// selector 查询条件
// error 错误信息
func (mc *client) RemoveOne(collection string, selector interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		_, err := c.DeleteOne(context.Background(), selector)
		return err
	})
}

// RemoveAll 根据条件删除所有符合条件的记录
// collection 集合名称
// selector 查询条件
// error 错误信息
func (mc *client) RemoveAll(collection string, selector interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		_, err := c.DeleteMany(context.Background(), selector)
		return err
	})
}

// CountById 根据id统计记录条数
// collection 集合名称
// id id
// n 总条数
func (mc *client) CountById(collection string, id interface{}) (n int64) {
	_ = mc.WithC(collection, func(c *mongo.Collection) error {
		var err error
		n, err = c.CountDocuments(context.Background(), bson.M{"_id": id})
		return err
	})
	return n
}

// Count 根据查询条件统计记录条数
// collection 集合名称
// query 查询条件
// n 总条数
func (mc *client) Count(collection string, query interface{}) (n int64) {
	_ = mc.WithC(collection, func(c *mongo.Collection) error {
		var err error
		n, err = c.CountDocuments(context.Background(), query)
		return err
	})
	return n
}

// Exist 根据查询条件判断是否存在
// collection 集合的名称
// query 查询条件
// bool true-存在，false-不存在
func (mc *client) Exist(collection string, query interface{}) bool {
	return mc.Count(collection, query) != 0
}

// ExistById 根据id判断是否存在
// collection 集合的名称
// id
// bool true-存在，false-不存在
func (mc *client) ExistById(collection string, id interface{}) bool {
	return mc.CountById(collection, id) != 0
}

// Page 根据查询条件，分页获取查询结果
// collection 集合名称
// query 查询条件
// offset 需要忽略的条数
// limit 每页的大小
// result 查询结果
// error 错误信息
func (mc *client) Page(collection string, query interface{}, offset int64, limit int64, result interface{}) error {
	return mc.WithC(collection, func(c *mongo.Collection) error {
		cursor, err := c.Find(context.Background(), query, options.Find().SetLimit(limit).SetSkip(offset))
		if nil != err {
			return err
		}

		return cursor.All(context.Background(), result)
	})
}

// PageAndCount 根据查询条件，分页获取查询结果，并且获取所有符合条件的条数
// collection 集合名称
// query 查询条件
// offset 需要忽略的条数
// limit 每页的大小
// result 查询结果
// total 总条数
// err 错误信息
func (mc *client) PageAndCount(collection string, query interface{}, offset int64, limit int64, result interface{}) (total int64, err error) {
	err = mc.WithC(collection, func(c *mongo.Collection) error {
		total, err = c.CountDocuments(context.Background(), query)
		if err != nil {
			return err
		}
		cursor, err := c.Find(context.Background(), query, options.Find().SetLimit(limit).SetSkip(offset))
		if nil != err {
			return err
		}

		return cursor.All(context.Background(), result)
	})
	return total, err
}

// PageSortAndCount 根据查询条件，分页获取指定排序的结果，并且获取所有符合条件的条数
// collection 集合名称
// query 查询条件
// sort 排序字段
// offset 需要忽略的条数
// limit 每页的大小
// result 查询结果
// total 总条数
// err 错误信息
func (mc *client) PageSortAndCount(collection string, query, sort interface{}, offset int64, limit int64, result interface{}) (total int64, err error) {
	err = mc.WithC(collection, func(c *mongo.Collection) error {
		total, err = c.CountDocuments(context.Background(), query)
		if err != nil {
			return err
		}
		cursor, err := c.Find(context.Background(), query, options.Find().SetLimit(limit).SetSkip(offset).SetSort(sort))
		if nil != err {
			return err
		}

		return cursor.All(context.Background(), result)
	})
	return total, err
}

// SetById 根据id设置文档，当id不存在的时候不会报错
// collection 集合名称
// id id
// change 文档
// error 错误信息
func (mc *client) SetById(collection string, id interface{}, change interface{}) error {
	return mc.UpdateById(collection, id, bson.M{"$set": change})
}
