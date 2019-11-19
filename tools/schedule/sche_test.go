package schedule

import (
	log "github.com/sirupsen/logrus"
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
	log.SetLevel(log.DebugLevel)
	sche := NewSchedule()
	job1 := sche.Delay(1 * time.Second).Do(func1)
	job2 := sche.Delay(4 * time.Second).Do(func2)
	job3 := sche.Every(1 * time.Second).Do(func1)
	time.Sleep(5 * time.Second)
	sche.Cancel(job2)
	sche.Cancel(job3)
	sche.Cancel(job1)

	time.Sleep(20 * time.Second)
}
