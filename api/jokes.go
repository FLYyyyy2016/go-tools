package api

import (
	"encoding/json"
	"fmt"
	"log"
)


type JokeResponse struct {
	Return  int         `json:"ret"`
	Data    JokeData `json:"data"`
	Message string      `json:"msg"`
}
type JokeData struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Joke [][]string
}

func GetJokes(num int) ([][]string, error) {
	queryUrl := appKey + fmt.Sprint(num)+ JokeServiceName
	sign, err := getSign(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryUrl = myURL + "?" + "&app_key=" + appKey + "&service=" + JokeServiceName + "&num=" +fmt.Sprint(num)+ "&sign=" + sign
	log.Println(queryUrl, "request")
	body := getRequest(queryUrl)
	var jokeResponse JokeResponse
	err = json.Unmarshal(body, &jokeResponse)
	if err != nil {
		log.Fatalln(err)
	}
	if jokeResponse.Return != 200 {
		return jokeResponse.Data.Joke, &getMessageError{jokeResponse.Return, jokeResponse.Message}
	}
	return jokeResponse.Data.Joke, err
}
