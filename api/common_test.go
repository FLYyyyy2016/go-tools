package api

import (
	"encoding/json"
	"log"
	"net/url"
	"testing"
)

func TestGetRequest(t *testing.T) {
	bs := getRequest("http://hb9.api.yesapi.cn/?service=App.Common_Weather.LiveWeather&city=%E6%AD%A6%E6%B1%89&app_key=FDA8A4A24DD86227286B58D0F909EA29&sign=393399E6AF54A71F5C302A699ECB4F5E")
	var reply WeatherResponse
	_ = json.Unmarshal(bs, &reply)
	log.Printf("%+v", reply)
}

func TestSign(t *testing.T) {
	s, err := getSign("liufei")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
}

func TestNumberToAscii(t *testing.T) {
	i := 5
	s := string(i)
	u := url.QueryEscape(s)
	log.Println(i, s, u)
	log.Printf("%d", i)
}
