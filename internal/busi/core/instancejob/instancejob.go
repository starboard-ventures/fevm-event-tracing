package instancejob

import (
	"context"
	"errors"
	"event-trace/internal/busi/core/instancejob/dealproposal"
	"event-trace/internal/busi/core/instancejob/wfil"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/api"
	log "github.com/sirupsen/logrus"
)

var _ ExecuteJobFn = dealproposal.DealProposalCreate{}
var _ ExecuteJobFn = wfil.Wfil{}

type ExecuteJobFn interface {
	EventTracing(context.Context, *api.FullNodeStruct, ...string) error
	GetEventName() string
}

var cronJob = CronJob{
	jobs: make(map[string]*singletonJob),
}

type singletonJob struct {
	jobIsRunning bool
	node         *api.FullNodeStruct

	startTime time.Time
	endTime   time.Time

	minHeight uint64
	maxHeight uint64

	executeJobFn ExecuteJobFn
}

type CronJob struct {
	node *api.FullNodeStruct
	jobs map[string]*singletonJob
}

func NewCronJob(node *api.FullNodeStruct, minHeight, maxHeight uint64, fn ExecuteJobFn) *singletonJob {
	eventName := fn.GetEventName()

	if _, ok := cronJob.jobs[eventName]; !ok {
		cronJob.jobs[eventName] = &singletonJob{
			node:         node,
			minHeight:    minHeight,
			maxHeight:    maxHeight,
			executeJobFn: fn,
		}
	}
	return cronJob.jobs[eventName]
}

func (j *singletonJob) TracingJobExecute(ctx context.Context, args ...string) error {
	if j.jobIsRunning {
		str := fmt.Sprintf("The previous job has begun at the time: %v, pls wait for it finishes or ctrl^c it.", j.startTime)
		log.Infof(str)
		return errors.New(str)
	} else {
		{
			j.jobIsRunning = true
			j.startTime = time.Now()

			log.Infof("Job runs at time: %v, from: %v - to: %v", j.startTime, j.minHeight, j.maxHeight)
		}

		defer func() {
			j.jobIsRunning = false
			j.endTime = time.Now()

			log.Infof("Job has been finished: %v", j.endTime)
		}()

		return j.executeJobFn.EventTracing(ctx, j.node, args...)
	}
}
