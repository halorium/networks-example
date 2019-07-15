package worker

type Worker struct {
	Queues map[string]Queue
}

func New() *Worker {
	return &Worker{
		Queues: make(map[string]Queue),
	}
}
