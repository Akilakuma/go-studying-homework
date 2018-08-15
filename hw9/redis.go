package main

import (
	"encoding/json"
	"log"

	"github.com/gomodule/redigo/redis"
)

// getRedisWetherData 取得redis 天氣資料
func getRedisWetherData() WeatherInfo {

	// 打開redis
	redisConnect, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err.Error())
	}
	defer redisConnect.Close()

	// 取得天氣資料
	data, getErr := redis.Bytes(redisConnect.Do("HGET", "api:weather:data", "taipei"))
	if getErr != nil {
		log.Println(getErr.Error())
	}

	var weatherInfo WeatherInfo

	if data != nil {
		decodeErr := json.Unmarshal(data, &weatherInfo)
		if decodeErr != nil {
			log.Println(decodeErr.Error())
		}
	}

	return weatherInfo
}

// setRedisWetherData 取得redis 天氣資料
func setRedisWetherData(data string) {

	// 打開redis
	redisConnect, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err.Error())
	}
	defer redisConnect.Close()

	log.Println("set redis")
	// 取得天氣資料
	_, setErr := redisConnect.Do("HSET", "api:weather:data", "taipei", data)
	if setErr != nil {
		log.Println(setErr.Error())
	}
}
