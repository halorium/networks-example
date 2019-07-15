package httpclient

import "sync"

var pool []*Client

var mutex = &sync.Mutex{}

func FromPool() *Client {
	if len(pool) == 0 {
		return New()
	}

	var client *Client

	mutex.Lock()
	client, pool = pool[0], pool[1:] // shift
	mutex.Unlock()

	return client
}

func ToPool(client *Client) {
	mutex.Lock()
	pool = append(pool, client) // push
	mutex.Unlock()
}
