package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// WeatherInfo 天氣API 資料格式
type WeatherInfo struct {
	CurrentTemperature string                   `json:"currentTemperature"`
	CwbUpdateTime      string                   `json:"cwbUpdateTime"`
	JSONUpdateTime     string                   `json:"jsonUpdateTime"`
	Information        string                   `json:"information"`
	CityName           string                   `json:"cityName"`
	Week               []OriginWeek             `json:"week"`
	TodayAndTomorrow   []OriginTodayAndTomorrow `json:"todayAndTomorrow"`
}

// OriginWeek 天氣API 資料格式原本week的部分
type OriginWeek struct {
	Day     string `json:"day"`
	Morning struct {
		Temperature string `json:"temperature"`
		Situation   string `json:"situation"`
		Img         string `json:"img"`
	} `json:"morning"`
	Night struct {
		Temperature string `json:"temperature"`
		Situation   string `json:"situation"`
		Img         string `json:"img"`
	} `json:"night"`
}

// WeekDetail 天氣API 資料格式原本week內部struct的部分
type WeekDetail struct {
	Temperature string `json:"temperature"`
	Situation   string `json:"situation"`
	Img         string `json:"img"`
}

// OriginTodayAndTomorrow 天氣API 資料格式原本day and tomorrow的部分
type OriginTodayAndTomorrow struct {
	Day         string `json:"day"`
	Temperature string `json:"temperature"`
	Img         string `json:"img"`
	Situation   string `json:"situation"`
	Desc        string `json:"desc"`
	Rain        string `json:"rain"`
}

func main() {
	// 執行server
	server()

}

// server 啟動server
func server() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/weather", getWeather)
	e.Start(":1111")
}

// getWeather 天氣的api
func getWeather(c echo.Context) error {

	redisData := getRedisWetherData()

	// 如果沒有redis資料
	if redisData.CityName == "" {
		log.Println("redis 取無資料")
		dbData := getDBWetherData()
		// 如果沒有DB資料
		if dbData.CityName == "" {
			log.Println("db 取無資料")
			apiData := setWeatherAPIRequest()
			setRedisWetherData(string(apiData))
			setDBWetherData(apiData)

			log.Println("使用API資料")
			var weather WeatherInfo
			decodeErr := json.Unmarshal(apiData, &weather)
			if decodeErr != nil {
				log.Println(decodeErr.Error())
			}
			return c.JSON(http.StatusOK, weather)
		}
		log.Println("使用db資料")
		return c.JSON(http.StatusOK, dbData)
	}
	log.Println("使用redis資料")
	return c.JSON(http.StatusOK, redisData)
}
