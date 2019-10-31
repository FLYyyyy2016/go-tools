package api

import (
	"log"
	"testing"
)

func TestQueryByCity(t *testing.T) {
	weather, err := QueryByCity("北京")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", weather)

}
