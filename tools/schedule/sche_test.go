package schedule

import (
	"github.com/go-playground/assert/v2"
	"sync"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {

}

func TestSchedule_Delay(t *testing.T) {
	sche := NewSchedule()
	lock := sync.Mutex{}
	i := 0

	f := func() {
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(1 * time.Millisecond)
		i = i + 1
	}

	for i := 0; i < 100; i++ {
		sche.Delay(time.Duration(time.Duration(1000+i*10) * time.Millisecond)).Do(f)
	}
	time.Sleep(3 * time.Second)
	assert.Equal(t, i, 100)
}

func TestSchedule_Every(t *testing.T) {
	sche := NewSchedule()
	lock := sync.Mutex{}
	i := 0

	f := func() {
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(1 * time.Millisecond)
		i = i + 1
	}
	sche.Every(100 * time.Millisecond).Do(f)
	time.Sleep(4550 * time.Millisecond)
	assert.Equal(t, i, 45)
}

func TestDelayJob_Cancel(t *testing.T) {
	delays := []string{}
	sche := NewSchedule()
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
		jobid := sche.Delay(time.Duration(time.Duration(1000+i*10) * time.Millisecond)).Do(f)
		delays = append(delays, jobid)
	}
	time.Sleep(1500 * time.Millisecond)
	for _, delay := range delays {
		err := sche.Cancel(delay)
		if err != nil {
			temp++
		}
	}
	time.Sleep(1 * time.Second)
	assert.Equal(t, temp, i)
}

func TestEveryJob_Cancel(t *testing.T) {
	everys := []string{}
	sche := NewSchedule()
	lock := sync.Mutex{}
	i := 0
	f := func() {
		lock.Lock()
		defer lock.Unlock()
		time.Sleep(1 * time.Millisecond)
		i = i + 1
	}
	for i := 0; i < 10; i++ {
		jobid := sche.Every(100 * time.Millisecond).Do(f)
		everys = append(everys, jobid)
	}
	time.Sleep(1510 * time.Millisecond)
	for _, every := range everys {
		err := sche.Cancel(every)
		if err != nil {
			t.Log(err)
		}
	}
	time.Sleep(1 * time.Second)
	assert.Equal(t, i, 150)
}
