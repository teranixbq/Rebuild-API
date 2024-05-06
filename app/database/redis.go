package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"recything/app/config"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.AppConfig) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RDB_ADDR,
		Username: cfg.RDB_USER,
		Password: cfg.RDB_PASS,
	})

	return client
}

type Redis struct {
	r *redis.Client
}

func NewRedis(r *redis.Client) *Redis {
	return &Redis{
		r: r,
	}
}

func (rdb *Redis) SetString(key string, value interface{}) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = rdb.r.Set(ctx, key, v, 0).Err()
	if err != nil {
		return err
	}

	return err
}

func (rdb *Redis) GetString(key string, value interface{}) error {
	ctx := context.Background()
	v, err := rdb.r.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	if reflect.TypeOf(value).Kind() == reflect.Struct {
		if reflect.TypeOf(value).Kind() != reflect.Ptr {
			err := ErrMsg("data must be of type pointer if struct")
			log.SetFlags(0)
			log.Println(err)
		}
	}

	err = json.Unmarshal([]byte(v), value)
	if err != nil {
		return err
	}

	return nil
}

func (rdb *Redis) SetExString(key string, value interface{}, exp time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = rdb.r.SetEx(ctx, key, v, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rdb *Redis) DelString(key string) error {
	ctx := context.Background()
	err := rdb.r.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rdb *Redis) SetDelString(key string, value interface{}) error {
	errDel := rdb.DelString(key)
	if errDel != nil {
		errSet := rdb.SetString(key, value)
		if errSet != nil {
			return errSet
		}
	}

	errSet := rdb.SetString(key, value)
	if errSet != nil {
		return errSet
	}

	return nil
}

func ErrMsg(err string) string {
	const redColor = "\033[31m"
	const resetColor = "\033[0m"

	response := fmt.Sprintf("%serror%s: %s", redColor, resetColor, err)
	return response
}
