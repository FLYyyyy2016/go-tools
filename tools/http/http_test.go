package http

import (
	"io"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

var url string
var writer io.Writer

func init() {
	url = "http://localhost:8888/"
	var err error
	writer, err = os.OpenFile("/dev/null", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("error to open log")
	}
}
func TestMain(m *testing.M) {
	go server()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGetRequest(t *testing.T) {
	getResponse(url, writer)
}

func BenchmarkGetRequestByWg(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			getResponse(url, writer)
			wg.Done()
		}()
	}
	wg.Wait()
}
func BenchmarkGetRequestByNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getResponse(url, writer)
	}
}
func BenchmarkGetRequestByGoChanel(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	ch := make(chan int, 256)
	for i := 0; i < b.N; i++ {
		ch <- 1
		go func() {
			getResponse(url, writer)
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}
func TestGetRequestByGoChanel(t *testing.T) {
	a := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(122)
	ch := make(chan int, 256)
	for i := 0; i < 100; i++ {
		go func() {
			ch <- 1
			getResponse(url, writer)
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println(time.Now().Sub(a))
}

func TestGetRequestByNormal(t *testing.T) {
	a := time.Now()
	//wg := sync.WaitGroup{}
	//wg.Add(10000)
	//ch:=make(chan int,256)
	for i := 0; i < 100; i++ {
		//ch<-1
		getResponse(url, writer)
		//<-ch
		//wg.Done()
	}
	//wg.Wait()
	log.Println(time.Now().Sub(a))

}
