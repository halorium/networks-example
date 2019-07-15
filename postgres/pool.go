package postgres

var client *Client

func ClientFromPool() *Client {
	if client == nil {
		client = new()
	}

	return client
}

func ClientToPool(client *Client) {
}
