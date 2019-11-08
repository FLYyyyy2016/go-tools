package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type Weather struct {
	Date       string
	Time       string
	City       string
	Visibility string
	Weather    string
	Tem        string
	Win        string
	WinSpeed   int
	WinMeter   string
	Humidity   string
	Pressure   string
	Air        string
	AirPm25    string `json:"air_pm25"`
	AirLevel   string `json:"air_level"`
	AirTips    string `json:"air_tips"`
	Alarm      Alarm
}

type Alarm struct {
	AlarmType    string
	AlarmLevel   string
	AlarmContent string
}

type WeatherResponse struct {
	Return  int         `json:"ret"`
	Data    WeatherData `json:"data"`
	Message string      `json:"msg"`
}
type WeatherData struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Weather Weather
}

func QueryByCity(city string) (Weather, error) {
	queryUrl := appKey + city + WeatherServiceName
	sign, err := getSign(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryUrl = myURL + "?" + "&app_key=" + appKey + "&service=" + WeatherServiceName + "&city=" + url.QueryEscape(city) + "&sign=" + sign
	//log.Println(queryUrl, "request")
	body := getRequest(queryUrl)
	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		log.Fatalln(err, city)
	}
	if weatherResponse.Return != 200 {
		return weatherResponse.Data.Weather, &getMessageError{weatherResponse.Return, weatherResponse.Message}
	}
	return weatherResponse.Data.Weather, err
}

func (w Weather) String() string {
	return fmt.Sprintf(`

时间：%s
地点：%s
天气状况：%s
风向：%s
风速：%d
可视距离：%s


`, w.Time, w.City, w.Weather, w.Win, w.WinSpeed, w.Visibility)
}
