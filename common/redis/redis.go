package redisModel

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"

	"time"
)

func InitRedis(Addr, PassWord string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: PassWord,
		DB:       0,
	})
}

func (rs *RedisStore) GetRedisValue() (string, error) {
	GetKey := rs.Conn.Get(rs.Context, rs.PreKey)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}
	return GetKey.Val(), nil
}

func (rs *RedisStore) SetRedisValue(value string) error {
	SetKV := rs.Conn.Set(rs.Context, rs.PreKey, value, time.Minute*2)
	return SetKV.Err()
}

//用于页面缓存
type RedisStore struct {
	PreKey     string
	Conn       *redis.Client
	Expiration time.Duration //过期时间
	Context    context.Context
}

func NewRedisStore(preKey string, Conn *redis.Client, expire time.Duration, context context.Context) *RedisStore {
	return &RedisStore{
		PreKey:     preKey,
		Expiration: expire,
		Conn:       Conn,
		Context:    context,
	}
}

func (rs *RedisStore) SetRedisPages(val interface{}) error {
	bytes, _ := json.Marshal(val)
	SetKV := rs.Conn.Set(rs.Context, rs.PreKey, bytes, rs.Expiration)
	return SetKV.Err()
}

func (rs *RedisStore) GetRedisPages(key string, obj interface{}) bool {
	val, err := rs.Conn.Get(rs.Context, key).Result()
	if err != nil {
		return false
	} else {
		json.Unmarshal([]byte(val), &obj)
		return true
	}
}

//用于 点赞 关注等功能

type RedisSet struct {
	Id      int64
	Object  string
	Conn    *redis.Client
	Context context.Context
}

func NewRedisSet(context context.Context, Objet string, Id int64, Conn *redis.Client) *RedisSet {
	return &RedisSet{
		Id:      Id,
		Object:  Objet,
		Conn:    Conn,
		Context: context,
	}
}

//当前用户是否点赞 or 关注

func (rs *RedisSet) JudgeSetByUid() (bool, error) {
	val, err := rs.Conn.SIsMember(rs.Context, rs.Object, rs.Id).Result()
	if err != nil {
		return val, err
	}
	return val, err
}

//-

func (rs *RedisSet) SetAction() error {
	val, err := rs.Conn.SIsMember(rs.Context, rs.Object, rs.Id).Result()
	if err != nil {
		return err
	}
	if val == false {
		_, err := rs.Conn.SAdd(rs.Context, rs.Object, rs.Id).Result()
		if err != nil {
			return err
		}
	} else {
		_, err := rs.Conn.SRem(rs.Context, rs.Object, rs.Id).Result()
		if err != nil {
			return err
		}
	}

	return err
}

// 获赞次数

func (rs *RedisSet) SetNumberByUid() (int64, error) {
	val, err := rs.Conn.SCard(rs.Context, rs.Object).Result()
	if err != nil {
		return val, err
	}
	return val, err
}

func (rs *RedisSet) SetMembers() (int64, error) {
	val, err := rs.Conn.SCard(rs.Context, rs.Object).Result()
	if err != nil {
		return val, err
	}
	return val, err
}
