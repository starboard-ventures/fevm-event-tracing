package instancejob

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"busi/internal/busi/core/instancejob/dealproposal"

	"github.com/filecoin-project/lotus/api"
	log "github.com/sirupsen/logrus"
)

var insJob *job
var once sync.Once

type job struct {
	jobIsRunning bool
	startTime    time.Time
	endTime      time.Time

	executeJobFn dealproposal.DealProposalTracingCronFn
}

type CronJob struct {
	node *api.FullNodeStruct

	minHeight uint64
	maxHeight uint64

	*job
}

func NewCronJob(node *api.FullNodeStruct, minHeight, maxHeight uint64) *CronJob {
	cj := &CronJob{
		node: node,

		minHeight: minHeight,
		maxHeight: maxHeight,
	}

	once.Do(func() {
		insJob = &job{}
	})

	cj.job = insJob

	return cj
}

func (cj *CronJob) TracingJobExecute(ctx context.Context, fn dealproposal.DealProposalTracingCronFn) error {
	if cj.jobIsRunning {
		str := fmt.Sprintf("The previous job has begun at the time: %v, pls wait for it finishes or ctrl^c it.", cj.startTime)
		log.Infof(str)
		return errors.New(str)
	} else {
		{
			cj.jobIsRunning = true
			cj.startTime = time.Now()

			log.Infof("Job runs at time: %v, from: %v - to: %v", cj.startTime, cj.minHeight, cj.maxHeight)
		}

		defer func() {
			cj.jobIsRunning = false
			cj.endTime = time.Now()

			log.Infof("Job has been finished: %v", cj.endTime)
		}()

		return fn(ctx, cj.node)
	}
}
