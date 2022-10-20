package cron

import "github.com/robfig/cron"

type CronJob struct {
	cronJob *cron.Cron
}

func NewCronJob() *CronJob {
	return &CronJob{
		cronJob: cron.New(),
	}
}

func (c *CronJob) StartTask() error {
	var err error

	err = runSign(c.cronJob)

	if err != nil {
		return err
	}
	c.cronJob.Start()
	return nil
}

func (c *CronJob) StopTask() {
	c.cronJob.Stop()
}
