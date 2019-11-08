package schedule

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"
)

type Schedule struct {
	Tasks []Task
}

func NewSchedule() *Schedule {
	return &Schedule{Tasks: []Task{}}
}

func (sche *Schedule) Delay(duration time.Duration) *DelayJob {

	newJob := DelayJob{
		Job:      Job{nextId(), nil, make(chan struct{})},
		duration: duration,
	}
	sche.Tasks = append(sche.Tasks, &newJob)
	return &newJob
}

func (sche *Schedule) Every(duration time.Duration) *EveryJob {
	newJob := EveryJob{
		Job:      Job{nextId(), nil, make(chan struct{})},
		duration: duration,
	}
	sche.Tasks = append(sche.Tasks, &newJob)
	return &newJob
}

type Task interface {
	Do(func())
}

type DelayJob struct {
	Job
	duration time.Duration
}

type EveryJob struct {
	Job
	duration time.Duration
}

type Job struct {
	JobId string
	work  func()
	close chan struct{}
}

func (job *DelayJob) Do(f func()) {
	timer := time.NewTimer(job.duration)
	//defer timer.Stop()
	go func() {
		<-timer.C
		f()
	}()
}

func (job *EveryJob) Do(f func()) {
	timer := time.NewTicker(job.duration)
	//defer timer.Stop()
	go func() {
		for {
			select {
			case <-timer.C:
				f()
			case <-job.close:
				return
			}
		}
	}()
}

func nextId() string {
	m := md5.New()
	now := time.Now()
	timeBytes, err := now.MarshalBinary()
	if err != nil {
		log.Fatalln(err)
	}
	m.Write(timeBytes)
	bs := m.Sum(nil)
	return hex.EncodeToString(bs)
}
