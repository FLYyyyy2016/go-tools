package schedule

import (
	"sync"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestSchedule_Delay(t *testing.T) {
	schedule := NewSchedule()
	lock := sync.Mutex{}
	i := 0

	f := func() {
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(1 * time.Millisecond)
		i = i + 1
	}

	for i := 0; i < 100; i++ {
		schedule.Delay(time.Duration(1000+i*10) * time.Millisecond).Do(f)
	}
	time.Sleep(3 * time.Second)
	assert.Equal(t, i, 100)
}

func TestSchedule_Every(t *testing.T) {
	schedule := NewSchedule()
	lock := sync.Mutex{}
	i := 0

	f := func() {
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(1 * time.Millisecond)
		i = i + 1
	}
	schedule.Every(100 * time.Millisecond).Do(f)
	time.Sleep(4550 * time.Millisecond)
	assert.Equal(t, i, 45)
}

func TestDelayJob_Cancel(t *testing.T) {
	var delays []string
	schedule := NewSchedule()
	lock := sync.Mutex{}
	i := 0
	temp := 0
	f := func() {
		lock.Lock()
		defer lock.Unlock()
		//time.Sleep(1*time.Millisecond)
		i = i + 1
	}
	for i := 0; i < 100; i++ {
		jobID := schedule.Delay(time.Duration(1000+i*10) * time.Millisecond).Do(f)
		delays = append(delays, jobID)
	}
	time.Sleep(1500 * time.Millisecond)
	for _, delay := range delays {
		err := schedule.Cancel(delay)
		if err != nil {
			temp++
		}
	}
	time.Sleep(1 * time.Second)
	assert.Equal(t, temp, i)
}

func TestEveryJob_Cancel(t *testing.T) {
	var everyStrings []string
	schedule := NewSchedule()
	lock := sync.Mutex{}
	i := 0
	f := func() {
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(1 * time.Millisecond)
		i = i + 1
	}
	for i := 0; i < 10; i++ {
		jobID := schedule.Every(100 * time.Millisecond).Do(f)
		everyStrings = append(everyStrings, jobID)
	}
	time.Sleep(1510 * time.Millisecond)
	for _, every := range everyStrings {
		err := schedule.Cancel(every)
		if err != nil {
			t.Log(err)
		}
	}
	time.Sleep(1 * time.Second)
	assert.Equal(t, i, 150)
}
