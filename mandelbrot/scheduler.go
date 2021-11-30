package main

type Scheduler interface {
	Schedule(planner Planner)
}

type SerialScheduler struct{}

func (s SerialScheduler) Schedule(planner Planner) {
	for i := 0; i != planner.jobCount(); i++ {
		job := planner.getJob(i)
		job()
	}
}

type ParallelScheduler struct{}

func (s ParallelScheduler) Schedule(planner Planner) {
	jobCount := planner.jobCount()
	ch := make(chan int)

	for i := 0; i != jobCount; i++ {
		job := planner.getJob(i)

		go (func() {
			job()
			ch <- 1
		})()
	}

	doneCount := 0
	for doneCount < jobCount {
		<-ch
		doneCount++
	}
}
