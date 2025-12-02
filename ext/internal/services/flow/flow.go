package flow

import (
	"context"
	"time"
)

type task struct {
	id       string
	finished bool
	hasNext  bool
}

type Service struct {
	ctx    context.Context
	cancel context.CancelFunc
	tasks  chan *task
}

func NewFlow() *Service {
	ctx, cancel := context.WithCancel(context.Background())

	return &Service{
		ctx:    ctx,
		cancel: cancel,
		tasks:  make(chan *task),
	}
}

func (s *Service) Stop() {
	s.cancel()
}

func (s *Service) Usleep(taskId string, milliseconds int64) {
	go func() {
		select {
		case <-time.After(time.Duration(milliseconds) * time.Microsecond):
			s.tasks <- &task{
				id:       taskId,
				finished: true,
				hasNext:  false,
			}

			return
		case <-s.ctx.Done():
			return
		}
	}()
}
