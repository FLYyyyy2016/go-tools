package schedule

import (
	"log"
	"testing"
	"time"
)

func func1() {
	log.Println("hello world")
}

func func2() {
	log.Println("你好，世界")
}

func TestSchedule(t *testing.T) {
	sche := NewSchedule()
	sche.Delay(4 * time.Second).Do(func1)
	sche.Delay(2 * time.Second).Do(func2)
	sche.Every(4 * time.Second).Do(func1)

	time.Sleep(20 * time.Second)
}
