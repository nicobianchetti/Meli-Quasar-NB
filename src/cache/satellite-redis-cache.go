package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/model"
)

type redisCache struct {
	host    string
	db      int           // es un indice entre 0 y 15
	expires time.Duration //tiempo de expiracion
}

//NewRedisCache .
func NewRedisCache(host string, db int, exp time.Duration) interfaces.ISatelliteCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *model.DTORequestSatellites) error {

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

func (cache *redisCache) Get(key string) (*model.DTORequestSatellites, error) {
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

func (cache *redisCache) Delete(key string) error {
	client := cache.getClient()

	err := client.Del(key).Err()

	if err != nil {
		return err
	}

	return nil

}
