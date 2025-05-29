package worker

import (
	"context"
	"go.uber.org/fx"
	"log"
)

var Module = fx.Module("worker", fx.Provide(
	newTaskQueue,
),
	fx.Invoke(startWorker),
)

type WriteTask struct {
	ID   string
	Data []byte
}

type TaskQueue chan WriteTask

func newTaskQueue() TaskQueue {
	return make(chan WriteTask, 1024) // 带 buffer，可配置
}

func startWorker(lc fx.Lifecycle, queue TaskQueue) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				for task := range queue {
					log.Printf("Worker writing: %s", task.ID)
					//if err := writeToDisk(task.ID, task.Data); err != nil {
					//	log.Printf("Write error: %v", err)
					//} else {
					//	log.Printf("Write success: %s", task.ID)
					//}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Worker shutting down")
			close(queue)
			return nil
		},
	})
}
