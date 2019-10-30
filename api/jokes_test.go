package api

import (
	"log"
	"net/url"
	"testing"
)

func TestGetJokes(t *testing.T) {
	i:=5
	jokes,err:=GetJokes(i)
	if err != nil {
		log.Fatalln(err)
	}
	jokeArray :=jokes[0]
	if len(jokeArray)!=i{
		log.Fatalln("没有得到指定数量的joke")
	}
	for _,joke:=range jokeArray{
		log.Println(joke)
	}
}

func TestNumberToString(t *testing.T) {
	i:=5
	s:=string(i)
	u:=url.QueryEscape(s)
	log.Println(i,s,u)
	log.Printf("%d",i)
}