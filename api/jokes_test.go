package api

import (
	"log"
	"testing"
)

func TestGetJokes(t *testing.T) {
	i := 1000
	jokes, err := GetJokes(i)
	if err != nil {
		log.Fatalln(err)
	}
	jokeArray := jokes[0]
	if len(jokeArray) != i {
		log.Fatalln("没有得到指定数量的joke")
	}
	for _, joke := range jokeArray {
		log.Println(joke)
	}
}
