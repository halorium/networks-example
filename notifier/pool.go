package notifier

import "sync"

var pool []*Notifier

var mutex = &sync.Mutex{}

func FromPool() *Notifier {
	if len(pool) == 0 {
		return new()
	}

	var notifier *Notifier

	mutex.Lock()
	notifier, pool = pool[0], pool[1:] // shift
	mutex.Unlock()

	return notifier
}

func ToPool(notifier *Notifier) {
	mutex.Lock()
	pool = append(pool, notifier) // push
	mutex.Unlock()
}
