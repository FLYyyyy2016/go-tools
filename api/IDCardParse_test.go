package api

import (
	"log"
	"testing"
)

func TestGetIDInfoByID(t *testing.T) {
	//通过网络生成
	id := "513436200010318889"
	name := "扬圣"
	info, err := GetIDInfoByID(id)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)
	log.Printf("%+v", info)
}
