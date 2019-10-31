package api

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	myURL  = "http://hb9.api.yesapi.cn/"
	appKey = "FDA8A4A24DD86227286B58D0F909EA29"
	appSec = "Og28qy5569N09bDURBjcd4zHT6Ck8pYvkgm6ZASG1IoFkvWPaHfA1yu1e5yQEiL6Wu"

	WeatherServiceName = "App.Common_Weather.LiveWeather"
	JokeServiceName    = "App.Common_Joke.RandOne"
	IDCardServiceName  = "App.Common_IDCard.Parse"
)

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
