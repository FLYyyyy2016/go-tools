package schedule

import (
	"crypto/md5"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	jobPrepare jobStatus = iota
	jobRunning
	jobFinish
	jobCreating
	jobCancel
	jobNotExist
)

type Schedule struct {
	Tasks map[string]Task
}

type jobStatus uint

type jobIsRunningError struct{}

func (error jobIsRunningError) Error() string {
	return "job is running ,you can try it again"
}

type jobIsCreatingError struct{}

func (error jobIsCreatingError) Error() string {
	return "job is Creating ,you can try it again"
}

type jobIsCancelError struct{}

func (error jobIsCancelError) Error() string {
	return "job is canceled "
}

type jobIsFinishError struct{}

func (error jobIsFinishError) Error() string {
	return "job is finished"
}

type jobNotExistError struct{}

func (error jobNotExistError) Error() string {
	return "job not exist ,you can try it again"
}

func NewSchedule() *Schedule {
	return &Schedule{Tasks: make(map[string]Task)}
}

func (sche *Schedule) Delay(duration time.Duration) *DelayJob {
	newJob := DelayJob{
		Job:      Job{JobId: nextId(), close: make(chan struct{})},
		duration: duration,
	}
	newJob.status = jobCreating
	sche.Tasks[newJob.JobId] = &newJob
	return &newJob
}

func (sche *Schedule) Every(duration time.Duration) *EveryJob {
	newJob := EveryJob{
		Job:      Job{JobId: nextId(), close: make(chan struct{})},
		duration: duration,
	}
	newJob.status = jobCreating
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
	sync.Mutex
	JobId  string
	status jobStatus
	work   func()
	close  chan struct{}
}

func (job *Job) setStatus(status jobStatus) {
	job.Lock()
	defer job.Unlock()
	job.status = status
}
func (job *Job) getStatus() jobStatus {
	job.Lock()
	defer job.Unlock()
	return job.status
}

func (job *DelayJob) Cancel() error {
	status := job.getStatus()
	if status == jobPrepare {
		job.close <- struct{}{}
	} else if status == jobRunning {
		return jobIsRunningError{}
	} else if status == jobCancel {
		return jobIsCancelError{}
	} else if status == jobCreating {
		return jobIsCreatingError{}
	} else if status == jobFinish {
		return jobIsFinishError{}
	}
	return nil
}

func (job *EveryJob) Cancel() error {
	status := job.getStatus()
	if status == jobPrepare {
		job.close <- struct{}{}
	} else if status == jobRunning {
		return jobIsRunningError{}
	} else if status == jobCancel {
		return jobIsCancelError{}
	} else if status == jobCreating {
		return jobIsCreatingError{}
	} else if status == jobFinish {
		return jobIsFinishError{}
	}
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
	go func() {
		job.setStatus(jobPrepare)
		select {
		case <-timer.C:
			job.setStatus(jobRunning)
			f()
			job.setStatus(jobFinish)
		case <-job.close:
			log.Debugln("job", job.JobId, "close")
			timer.Stop()
			job.setStatus(jobCancel)
			return
		}
	}()
	return job.JobId
}

func (job *EveryJob) Do(f func()) string {
	timer := time.NewTicker(job.duration)
	go func() {
		job.setStatus(jobPrepare)
		for {
			select {
			case <-timer.C:
				job.setStatus(jobRunning)
				f()
				job.setStatus(jobPrepare)
			case <-job.close:
				log.Debugln("job", job.JobId, "close")
				timer.Stop()
				job.setStatus(jobCancel)
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
