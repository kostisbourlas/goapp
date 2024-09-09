package watcher

type Counter struct {
	Iteration int    `json:"iteration"`
	HexValue  string `json:"value"`
}

type CounterReset struct{}
