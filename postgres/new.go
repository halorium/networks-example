package postgres

func new() *Client {
	pg := open()

	ping(pg)

	return &Client{pg}
}
