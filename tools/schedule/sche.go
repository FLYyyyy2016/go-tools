package schedule

import (
	"crypto/md5"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"time"
)

type Schedule struct {
	Tasks map[string]Task
}

func NewSchedule() *Schedule {
	return &Schedule{Tasks: make(map[string]Task)}
}

func (sche *Schedule) Delay(duration time.Duration) *DelayJob {
	newJob := DelayJob{
		Job:      Job{nextId(), nil, make(chan struct{})},
		duration: duration,
	}
	sche.Tasks[newJob.JobId] = &newJob
	return &newJob
}

func (sche *Schedule) Every(duration time.Duration) *EveryJob {
	newJob := EveryJob{
		Job:      Job{nextId(), nil, make(chan struct{})},
		duration: duration,
	}
	sche.Tasks[newJob.JobId] = &newJob
	return &newJob
}

func (sche *Schedule) Cancel(jobId string) error {
	return sche.Tasks[jobId].Cancel()
}

type Task interface {
	Do(func()) string
	GetId() string
	Cancel() error
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

func (job *DelayJob) Cancel() error {
	job.close <- struct{}{}
	return nil
}

func (job *EveryJob) Cancel() error {
	job.close <- struct{}{}
	return nil
}

func (job *DelayJob) GetId() string {
	return job.JobId
}

func (job *EveryJob) GetId() string {
	return job.JobId
}

func (job *DelayJob) Do(f func()) string {
	timer := time.NewTimer(job.duration)
	//defer timer.Stop()
	go func() {
		select {
		case <-timer.C:
			f()
		case <-job.close:
			log.Debugln("job", job.JobId, "close")
			timer.Stop()
			return
		}
	}()
	return job.JobId
}

func (job *EveryJob) Do(f func()) string {
	timer := time.NewTicker(job.duration)
	//defer timer.Stop()
	go func() {
		for {
			select {
			case <-timer.C:
				f()
			case <-job.close:
				log.Debugln("job", job.JobId, "close")
				timer.Stop()
				return
			}
		}
	}()
	return job.JobId
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
