package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

//RedisCache .
type RedisCache struct {
	host    string
	db      int           // es un indice entre 0 y 15
	expires time.Duration //tiempo de expiracion
	pass    string
}

//NewRedisCache .
func NewRedisCache(host string, db int, exp time.Duration, pass string) interfaces.ISatelliteCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
		pass:    pass,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: cache.pass,
		DB:       cache.db,
	})
}

//Set .
func (cache *RedisCache) Set(key string, value *model.DTORequestSatellites) error {

	client := cache.getClient()

	// pong, err := client.Ping(client.Context()).Result()
	// fmt.Println(pong, err)

	json, err := json.Marshal(value)

	if err != nil {
		return err
	}

	err = client.Set(key, json, cache.expires*time.Second).Err()

	if err != nil {
		return err
	}

	return nil

}

//Get .
func (cache *RedisCache) Get(key string) (*model.DTORequestSatellites, error) {
	client := cache.getClient()

	val, err := client.Get(key).Result()

	//Si estructura viene vacía no lo considero un error,devuelvo nil
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	satellite := model.DTORequestSatellites{}
	err = json.Unmarshal([]byte(val), &satellite)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &satellite, nil
}

//Delete .
func (cache *RedisCache) Delete(key string) error {
	client := cache.getClient()

	err := client.Del(key).Err()

	if err != nil {
		return err
	}

	return nil

}
