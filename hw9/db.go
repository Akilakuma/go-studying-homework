package main

import (
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// WeatherHeader 天氣基本資料
type WeatherHeader struct {
	ID                 int    `gorm:"primary_key;column:id"`
	CurrentTemperature string `gorm:"column:current_temperature"`
	CwbUpdateTime      string `gorm:"column:cwb_update_time"`
	JSONUpdateTime     string `gorm:"column:json_update_time"`
	Information        string `gorm:"column:information"`
	CityName           string `gorm:"column:city_name"`
}

// Week 天氣當週資料
type Week struct {
	ID                 int    `gorm:"primary_key;column:id"`
	HeaderID           int    `gorm:"column:header_id"`
	Day                string `gorm:"column:day"`
	MorningTemperature string `gorm:"column:morning_temperature"`
	MorningSituation   string `gorm:"column:morning_situation"`
	MorningImg         string `gorm:"column:morning_img"`
	NightTemperature   string `gorm:"column:night_temperature"`
	NightSituation     string `gorm:"column:nitght_situation"`
	NightImg           string `gorm:"column:night_img"`
}

// TodayAndTomorrow 今天和明天資料
type TodayAndTomorrow struct {
	ID          int    `gorm:"primary_key;column:id"`
	HeaderID    int    `gorm:"column:header_id"`
	Day         string `gorm:"column:day"`
	Temperature string `gorm:"column:temperature"`
	Img         string `gorm:"column:img"`
	Situation   string `gorm:"column:situation"`
	Desc        string `gorm:"column:desc"`
	Rain        string `gorm:"column:rain"`
}

// TableName rename table weather_header
func (WeatherHeader) TableName() string {
	return "weather_header"
}

// TableName rename table weather_week
func (Week) TableName() string {
	return "weather_week"
}

// TableName rename table weather_day_tomorrow
func (TodayAndTomorrow) TableName() string {
	return "weather_day_tomorrow"
}

// setDBWetherData 寫資料到DB
func setDBWetherData(data []byte) {
	// 建立連線
	// [root] 帳號，
	// [qwe123] 密碼
	// [tcp(127.0.0.1:3306)] ip和port而且一定要外刮一層tcp
	// [Match] DB的名稱
	db, err := gorm.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	// 選擇要不要打開debug mode 超實用
	// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	db.LogMode(false)

	// 連不連得到，這個超級重要!!
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	log.Println("set db")

	var weatherInfo WeatherInfo
	if data != nil {
		decodeErr := json.Unmarshal(data, &weatherInfo)
		if decodeErr != nil {
			log.Println(decodeErr.Error())
		}
	}

	header := WeatherHeader{
		CurrentTemperature: weatherInfo.CurrentTemperature,
		CwbUpdateTime:      weatherInfo.CwbUpdateTime,
		JSONUpdateTime:     weatherInfo.JSONUpdateTime,
		Information:        weatherInfo.Information,
		CityName:           weatherInfo.CityName,
	}

	db.Create(&header)
	headerID := header.ID

	for _, v := range weatherInfo.Week {
		db.Create(&Week{
			HeaderID:           headerID,
			Day:                v.Day,
			MorningTemperature: v.Morning.Temperature,
			MorningSituation:   v.Morning.Situation,
			MorningImg:         v.Morning.Img,
			NightTemperature:   v.Night.Temperature,
			NightSituation:     v.Night.Situation,
			NightImg:           v.Night.Img,
		})

	}

	for _, v := range weatherInfo.TodayAndTomorrow {
		db.Create(&TodayAndTomorrow{
			HeaderID:    headerID,
			Day:         v.Day,
			Temperature: v.Temperature,
			Img:         v.Img,
			Situation:   v.Situation,
			Desc:        v.Desc,
			Rain:        v.Rain,
		})
	}
}

// getDBWetherData 從DB讀資料出來
func getDBWetherData() WeatherInfo {
	db, err := gorm.Open("mysql", "root:qwe123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	// 選擇要不要打開debug mode 超實用
	// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	db.LogMode(false)

	// 連不連得到，這個超級重要!!
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	// 組header
	var weatherHeader WeatherHeader
	db.Table("weather_header").Where("weather_header.city_name = ?", "臺北市").Order("id desc").First(&weatherHeader)

	headerID := weatherHeader.ID

	// 組week
	var week []Week
	db.Table("weather_week").Where("header_id = ?", headerID).Find(&week)
	var apiWeekData []OriginWeek
	for _, v := range week {
		moring := WeekDetail{
			Temperature: v.MorningTemperature,
			Situation:   v.MorningSituation,
			Img:         v.MorningImg,
		}
		night := WeekDetail{
			Temperature: v.NightTemperature,
			Situation:   v.NightSituation,
			Img:         v.NightImg,
		}

		apiWeek := OriginWeek{
			Day:     v.Day,
			Morning: moring,
			Night:   night,
		}
		apiWeekData = append(apiWeekData, apiWeek)
	}

	// 組day and tomorrorw
	var day []OriginTodayAndTomorrow
	db.Table("weather_day_tomorrow").Where("header_id = ?", headerID).Find(&day)
	for _, v := range day {
		apiDay := OriginTodayAndTomorrow{
			Day:         v.Day,
			Temperature: v.Temperature,
			Img:         v.Img,
			Situation:   v.Situation,
			Desc:        v.Desc,
			Rain:        v.Rain,
		}
		day = append(day, apiDay)
	}

	return WeatherInfo{
		CurrentTemperature: weatherHeader.CurrentTemperature,
		CwbUpdateTime:      weatherHeader.CwbUpdateTime,
		JSONUpdateTime:     weatherHeader.JSONUpdateTime,
		Information:        weatherHeader.Information,
		CityName:           weatherHeader.CityName,
		Week:               apiWeekData,
		TodayAndTomorrow:   day,
	}
}
