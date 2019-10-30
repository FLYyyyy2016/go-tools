package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	myURL       = "http://hb9.api.yesapi.cn/"
	serviceName = "App.Common_Weather.LiveWeather"
	appKey      = "FDA8A4A24DD86227286B58D0F909EA29"
	appSec      = "Og28qy5569N09bDURBjcd4zHT6Ck8pYvkgm6ZASG1IoFkvWPaHfA1yu1e5yQEiL6Wu"
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
	queryUrl := appKey + city + serviceName
	sign, err := getSign(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryUrl = myURL + "?" + "&app_key=" + appKey + "&service=" + serviceName + "&city=" + url.QueryEscape(city) + "&sign=" + sign
	log.Println(queryUrl, "request")
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

func getRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("request", url, err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("request", url, err)
	}
	return bytes
}

func getSign(s string) (string, error) {
	s += appSec
	gen := md5.New()
	gen.Write([]byte(s))
	bs := gen.Sum(nil)
	str := fmt.Sprintf("%x", bs)
	if len(bs) != 16 {
		return "", &md5GetError{}
	}
	return strings.ToUpper(str), nil
}

type md5GetError struct{}

func (err *md5GetError) Error() string {
	return "get md5 sum error"
}

type getMessageError struct {
	code int
	msg  string
}

func (err *getMessageError) Error() string {
	return fmt.Sprintf("error code is %d \nerr message is %s", err.code, err.msg)
}
