package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getResponse(url string, writer io.Writer) {

	res, err := http.Get(url)
	if err != nil {
		log.Println("get error")
	}

	_, err = io.Copy(writer, io.MultiReader(res.Body, strings.NewReader("\n")))
	if err != nil {
		log.Println("get body error")
	}

	fmt.Fprintln(writer, res.ContentLength)
	fmt.Fprintln(writer, res.Proto)
}
